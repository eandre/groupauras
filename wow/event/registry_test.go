package event
import (
	"testing"
	"reflect"
)

func TestEvents(t *testing.T) {
	type event struct {
		event string
		args []interface{}
	}
	var received []event

	handler := func(e string, args ...interface{}) {
		received = append(received, event{e, args})
	}
	Register("foo", handler)
	Register("bar", handler)

	Trigger("baz")
	if len(received) != 0 {
		t.Errorf("Got %d events, want 0", len(received))
	}

	Trigger("foo")
	Trigger("bar", 1, 2, 3)
	want := []event{{"foo", nil}, {"bar", []interface{}{1, 2, 3}}}
	if !reflect.DeepEqual(received, want) {
		t.Errorf("Got events %+v, want %+v", received, want)
	}
}