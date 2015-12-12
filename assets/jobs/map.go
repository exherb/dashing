package jobs

import (
	// "fmt"
	"github.com/exherb/Dashing"
	"time"
)

type MapJob struct{}

func (j MapJob) Work(send chan *dashing.Event) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			send <- &dashing.Event{"map", "widget", "map", map[string]interface{}{
				"coordinate": []int{random(73, 135), random(3, 53)},
				"value":      random(0, 100),
			}}
		}
	}
}

func init() {
	dashing.Register(MapJob{})
}
