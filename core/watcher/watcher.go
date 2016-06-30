package watcher

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/core/library"
	"github.com/eandre/groupauras/core/runner"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type Watcher struct {
	all       map[string]*aura.CompiledAura // aura id -> compiled aura
	enables   map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool
	disables  map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool
	listeners map[string]int
	runner    *runner.Runner
}

func New(l *library.Library, r *runner.Runner) *Watcher {
	w := &Watcher{
		all:       make(map[string]*aura.CompiledAura),
		enables:   make(map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool),
		disables:  make(map[string]map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool),
		listeners: make(map[string]int),
		runner:    r,
	}
	l.RegisterListener(w)
	return w
}

func (w *Watcher) OnAuraChanged(a *aura.Aura, enabled bool) {
	// TODO need to gracefully handle changing of running auras
	prev := w.all[a.ID]
	if prev != nil {
		w.removeAura(prev)
	}
	if !enabled {
		return
	}

	ca, err := aura.Compile(a)
	if err != nil {
		println("Could not compile aura:", err.Error())
		return
	}
	w.addAura(ca)
}

func (w *Watcher) addAura(ca *aura.CompiledAura) {
	if w.all[ca.Aura.ID] != nil {
		return
	}
	w.all[ca.Aura.ID] = ca

	for event, f := range ca.Enables {
		m := w.enables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool)
			w.enables[event] = m
		}
		m[ca] = f
		w.addListener(event)
	}

	for event, f := range ca.Disables {
		m := w.disables[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]func(*aura.CompiledAura, string, []interface{}) bool)
			w.disables[event] = m
		}
		m[ca] = f
		w.addListener(event)
	}
}

func (w *Watcher) removeAura(ca *aura.CompiledAura) {
	if w.all[ca.Aura.ID] == nil {
		return
	}
	delete(w.all, ca.Aura.ID)

	for event := range ca.Enables {
		m := w.enables[event]
		delete(m, ca)
		if len(m) == 0 {
			delete(w.enables, event)
		}
		w.removeListener(event)
	}

	for event := range ca.Disables {
		m := w.disables[event]
		delete(m, ca)
		if len(m) == 0 {
			delete(w.disables, event)
		}
		w.removeListener(event)
	}
}

func (w *Watcher) addListener(event string) {
	n := w.listeners[event]
	w.listeners[event] = n + 1
	if n == 0 {
		wow.RegisterEvent(event, w.onEvent)
	}
}

func (w *Watcher) removeListener(event string) {
	n := w.listeners[event]
	w.listeners[event] = n - 1
	if n == 1 {
		delete(w.listeners, event)
		wow.UnregisterEvent(event, w.onEvent)
	}
}

func (w *Watcher) onEvent(event string, args []interface{}) {
	for ca, f := range w.enables[event] {
		if !w.runner.Running(ca) && f(ca, event, args) {
			w.runner.Start(ca)
		}
	}

	for ca, f := range w.disables[event] {
		if w.runner.Running(ca) && f(ca, event, args) {
			w.runner.Stop(ca)
		}
	}
}
