package trigger

import (
	"github.com/eandre/groupauras/aura"
	"github.com/eandre/groupauras/bridge"
)

type Watcher struct {
	all       map[*aura.CompiledAura]bool
	enables   map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool
	disables  map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool
	listeners map[string]int
}

func NewWatcher() *Watcher {
	return &Watcher{
		all:       make(map[*aura.CompiledAura]bool),
		enables:   make(map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool),
		disables:  make(map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool),
		listeners: make(map[string]int),
	}
}

func (t *Watcher) Add(a *aura.CompiledAura) {
	if t.all[a] {
		return
	}
	t.all[a] = true

	for event, f := range a.Enables {
		m := t.enables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool)
			t.enables[event] = m
		}
		m[a] = f
		t.addListener(event)
	}

	for event, f := range a.Disables {
		m := t.disables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool)
			t.disables[event] = m
		}
		m[a] = f
		t.addListener(event)
	}
}

func (t *Watcher) Remove(a *aura.CompiledAura) {
	if !t.all[a] {
		return
	}
	delete(t.all, a)

	for event := range a.Enables {
		m := t.enables[event]
		delete(m, a)
		if len(m) == 0 {
			delete(t.enables, event)
		}
		t.removeListener(event)
	}

	for event := range a.Disables {
		m := t.disables[event]
		delete(m, a)
		if len(m) == 0 {
			delete(t.disables, event)
		}
		t.removeListener(event)
	}
}

func (t *Watcher) addListener(event string) {
	n := t.listeners[event]
	t.listeners[event] = n + 1
	if n == 0 {
		bridge.RegisterEvent(event, t.onEvent)
	}
}

func (t *Watcher) removeListener(event string) {
	n := t.listeners[event]
	t.listeners[event] = n - 1
	if n == 1 {
		delete(t.listeners, event)
		bridge.UnregisterEvent(event, t.onEvent)
	}
}

func (t *Watcher) onEvent(event string, args []interface{}) {
	for aura, f := range t.enables[event] {
		if !aura.Active && f(aura, event, args) {
			aura.Activate()
		}
	}

	for aura, f := range t.disables[event] {
		if aura.Active && f(aura, event, args) {
			aura.Deactivate()
		}
	}
}
