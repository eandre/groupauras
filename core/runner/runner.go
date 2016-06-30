package runner

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type Runner struct {
	running   map[*aura.CompiledAura]bool
	updateMap map[*aura.CompiledAura]bool
	eventMap  map[string]map[*aura.CompiledAura]bool
}

func New() *Runner {
	return &Runner{
		running:   make(map[*aura.CompiledAura]bool),
		updateMap: make(map[*aura.CompiledAura]bool),
		eventMap:  make(map[string]map[*aura.CompiledAura]bool),
	}
}

func (r *Runner) Running(ca *aura.CompiledAura) bool {
	return r.running[ca]
}

func (r *Runner) Start(ca *aura.CompiledAura) {
	if r.running[ca] {
		return
	}
	r.running[ca] = true

	if ca.OnActivate != nil {
		ca.OnActivate(ca)
	}
	if ca.OnUpdate != nil {
		r.updateMap[ca] = true
		wow.RegisterUpdate(r.onUpdate)
	}
	for event := range ca.Events {
		m := r.eventMap[event]
		if m == nil {
			m = make(map[*aura.CompiledAura]bool)
			wow.RegisterEvent(event, r.onEvent)
		}
		m[ca] = true
	}
}

func (r *Runner) Stop(ca *aura.CompiledAura) {
	if !r.running[ca] {
		return
	}
	delete(r.running, ca)

	for event := range ca.Events {
		m := r.eventMap[event]
		delete(m, ca)
		if len(m) == 0 {
			wow.UnregisterEvent(event, r.onEvent)
			delete(r.eventMap, event)
		}
	}

	if ca.OnUpdate != nil {
		delete(r.updateMap, ca)
		if len(r.updateMap) == 0 {
			wow.UnregisterUpdate(r.onUpdate)
		}
	}

	if ca.OnDeactivate != nil {
		ca.OnDeactivate(ca)
	}
}

func (r *Runner) onEvent(event string, args []interface{}) {
	for ca := range r.eventMap[event] {
		ca.Events[event](ca, event, args)
	}
}

func (r *Runner) onUpdate(dt float32) {
	for ca := range r.updateMap {
		ca.OnUpdate(ca, dt)
	}
}
