package event
import "sync"

var (
	eventsMu sync.RWMutex
	events map[string][]Handler
)

func Register(event string, handler Handler) {
	eventsMu.Lock()
	defer eventsMu.Unlock()

	events[event] = append(events[event], handler)
}

func Trigger(event string, args ...interface{}) {
	eventsMu.RLock()
	defer eventsMu.RUnlock()

	for _, handler := range events[event] {
		handler(event, args...)
	}
}

type Handler func(event string, args ...interface{})

func init() {
	events = make(map[string][]Handler)
}
