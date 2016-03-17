package activity

import (
	"github.com/eandre/sbm/groupauras/aura"
	"github.com/eandre/sbm/groupauras/bridge"
)

type Tracker struct {
	all       map[*aura.CompiledAura]bool
	enables   map[string]map[*aura.CompiledAura]func(event string, args []interface{}) bool
	disables  map[string]map[*aura.CompiledAura]func(event string, args []interface{}) bool
	listeners map[string]int

	enableFunc  func(*aura.CompiledAura)
	disableFunc func(*aura.CompiledAura)
}

func NewTracker(enable, disable func(*aura.CompiledAura)) *Tracker {
	return &Tracker{
		all:         make(map[*aura.CompiledAura]bool),
		enables:     make(map[string]map[*aura.CompiledAura]func(event string, args []interface{}) bool),
		disables:    make(map[string]map[*aura.CompiledAura]func(event string, args []interface{}) bool),
		listeners:   make(map[string]int),
		enableFunc:  enable,
		disableFunc: disable,
	}
}

func (t *Tracker) Add(a *aura.CompiledAura) {
	if t.all[a] {
		return
	}
	t.all[a] = true

	for event, f := range a.Enables {
		m := t.enables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(event string, args []interface{}) bool)
			t.enables[event] = m
		}
		m[a] = f
		t.addListener(event)
	}

	for event, f := range a.Disables {
		m := t.disables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(event string, args []interface{}) bool)
			t.disables[event] = m
		}
		m[a] = f
		t.addListener(event)
	}
}

func (t *Tracker) Remove(a *aura.CompiledAura) {
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

func (t *Tracker) addListener(event string) {
	n := t.listeners[event]
	t.listeners[event] = n + 1
	if n == 0 {
		bridge.RegisterEvent(event, t.onEvent)
	}
}

func (t *Tracker) removeListener(event string) {
	n := t.listeners[event]
	t.listeners[event] = n - 1
	if n == 1 {
		delete(t.listeners, event)
		bridge.UnregisterEvent(event, t.onEvent)
	}
}

func (t *Tracker) onEvent(event string, args []interface{}) {
	for aura, f := range t.enables[event] {
		if f(event, args) {
			t.enableFunc(aura)
		}
	}

	for aura, f := range t.disables[event] {
		if f(event, args) {
			t.disableFunc(aura)
		}
	}
}

func countListeners(m map[string]map[*aura.CompiledAura]bool, event string) int {
	m2 := m[event]
	if m2 == nil {
		return 0
	}
	return len(m2)
}

type trackListener struct {
	tracker *Tracker
}
