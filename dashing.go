package dashing

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

var config *Config
var eventBroker *EventBroker

func renderStatus(w http.ResponseWriter, status int, key string, info string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{key: info})
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		dashboards, _ := filepath.Glob(filepath.Join(config.DashingPath, "dashboards", "*.json"))
		if len(dashboards) > 0 {
			defaultDashboard := config.DefaultDashboard
			if len(defaultDashboard) == 0 {
				firstDashboard := dashboards[0]
				firstDashboard = filepath.Base(firstDashboard)
				defaultDashboard = firstDashboard[0 : len(firstDashboard)-5]
			}
			http.Redirect(w, r, "/"+defaultDashboard, http.StatusMovedPermanently)
		} else {
			renderStatus(w, 404, "status", "not found")
		}
	}
}

func dashboard(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	layoutHTMLBytes := MustAsset("layout/dashboard.html")
	w.Write(layoutHTMLBytes)
}

func dashboardPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dashboard := params.ByName("dashboard")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		renderStatus(w, 400, "status", err.Error())
		return
	}

	token := data["token"]
	if token != config.GlobalToken {
		renderStatus(w, 403, "status", "forbidden")
		return
	}
	delete(data, "token")

	eventBroker.events <- &Event{dashboard, "dashboard", dashboard, data}
	renderStatus(w, 200, "status", "ok")
}

func widgetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dashboard := params.ByName("dashboard")
	widget := params.ByName("widget")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		renderStatus(w, 400, "status", err.Error())
		return
	}

	token := data["token"]
	if token != config.GlobalToken &&
		token != config.HackToken {
		widgetToken := config.WidgetTokens[widget]
		if len(widgetToken) == 0 || token != widgetToken {
			renderStatus(w, 403, "status", "forbidden")
			return
		}
	}

	delete(data, "token")
	eventBroker.events <- &Event{dashboard, "widget", widget, data}
	renderStatus(w, 200, "status", "ok")
}

func events(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dashboard := params.ByName("dashboard")
	f, ok := w.(http.Flusher)
	if !ok {
		renderStatus(w, 500, "status", "error: streaming unsupported")
		return
	}
	closeNotification, ok := w.(http.CloseNotifier)
	if !ok {
		renderStatus(w, 500, "status", "error: close notification unsupported")
		return
	}

	events := make(chan *Event)
	eventBroker.newClients <- events
	defer func() {
		eventBroker.defunctClients <- events
	}()

	w.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	closer := closeNotification.CloseNotify()
	for {
		select {
		case event := <-events:
			if event.Dashboard == dashboard || event.Dashboard == "*" {
				data := event.Body
				data["id"] = event.ID
				if event.Target != "" {
					w.Write([]byte("event: "))
					w.Write([]byte(event.Target))
					w.Write([]byte("\n"))
				}
				raw, _ := json.Marshal(data)
				w.Write([]byte("data: "))
				w.Write(raw)
				w.Write([]byte("\n\n"))
				f.Flush()
			}
		case <-closer:
			break
		}
	}
}

func Server(dashingPath string, host string, port int) error {
	config = NewConfig(filepath.Join(dashingPath, "config.json"))
	config.DashingPath = dashingPath

	eventBroker = NewEventBlock()
	eventBroker.Start()

	for _, j := range registry {
		go j.Work(eventBroker.events)
	}

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/:dashboard", dashboard)
	router.POST("/dashboards/:dashboard", dashboardPost)
	router.POST("/dashboards/:dashboard/widgets/:widget", widgetPost)
	router.GET("/:dashboard/events", events)
	router.NotFound = http.FileServer(http.Dir(filepath.Join(dashingPath, "public")))

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Listening on %s ...\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		return err
	}
	return nil
}
