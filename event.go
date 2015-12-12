package dashing

type Event struct {
	Dashboard string
	Target string
	ID     string
	Body   map[string]interface{}
}

type EventBroker struct {
	clients        map[chan *Event]bool
	newClients     chan chan *Event
	defunctClients chan chan *Event
	events         chan *Event
	lastEvents     map[string]*Event
}

func (b *EventBroker) Start() {
	go func() {
		for {
			select {
			case s := <-b.newClients:
				b.clients[s] = true
				for _, e := range b.lastEvents {
					s <- e
				}
			case s := <-b.defunctClients:
				delete(b.clients, s)
			case event := <-b.events:
				b.lastEvents[event.ID] = event
				for s := range b.clients {
					s <- event
				}
			}
		}
	}()
}

func NewEventBlock() *EventBroker {
	return &EventBroker{
		make(map[chan *Event]bool),
		make(chan (chan *Event)),
		make(chan (chan *Event)),
		make(chan *Event),
		map[string]*Event{},
	}
}
