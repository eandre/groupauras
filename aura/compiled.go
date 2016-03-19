package aura

import "github.com/eandre/groupauras/bridge"

type CompiledAura struct {
	Aura     *Aura
	Enables  map[string]func(event string, args []interface{}) bool
	Disables map[string]func(event string, args []interface{}) bool
	Active   bool

	// Relevant for activate auras
	events       map[string]func(aura *CompiledAura, event string, args []interface{})
	onUpdate     func(aura *CompiledAura, dt float32)
	onActivate   func(aura *CompiledAura)
	onDeactivate func(aura *CompiledAura)
}

func (ca *CompiledAura) Activate() {
	if ca.Active {
		return
	}
	ca.Active = true

	if ca.onActivate != nil {
		ca.onActivate(ca)
	}
	if ca.onUpdate != nil {
		bridge.RegisterUpdate(ca.doOnUpdate)
	}
	for event := range ca.events {
		bridge.RegisterEvent(event, ca.onEvent)
	}
}

func (ca *CompiledAura) Deactivate() {
	if !ca.Active {
		return
	}
	ca.Active = false

	for event := range ca.events {
		bridge.UnregisterEvent(event, ca.onEvent)
	}
	if ca.onUpdate != nil {
		bridge.UnregisterUpdate(ca.doOnUpdate)
	}
	if ca.onDeactivate != nil {
		ca.onDeactivate(ca)
	}
}

func (ca *CompiledAura) onEvent(event string, args []interface{}) {
	ca.events[event](ca, event, args)
}

func (ca *CompiledAura) doOnUpdate(dt float32) {
	ca.onUpdate(ca, dt)
}

func Compile(aura *Aura) (*CompiledAura, error) {
	ca := newCompiledAura(aura)

	for event, src := range aura.Enables {
		f, err := bridge.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.Enables[event] = f.(func(string, []interface{}) bool)
	}
	for event, src := range aura.Disables {
		f, err := bridge.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.Disables[event] = f.(func(string, []interface{}) bool)
	}
	for event, src := range aura.Events {
		f, err := bridge.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.events[event] = f.(func(*CompiledAura, string, []interface{}))
	}
	if aura.OnUpdate != "" {
		f, err := bridge.Eval(aura.OnUpdate)
		if err != nil {
			return nil, err
		}
		ca.onUpdate = f.(func(*CompiledAura, float32))
	}
	if aura.OnActivate != "" {
		f, err := bridge.Eval(aura.OnActivate)
		if err != nil {
			return nil, err
		}
		ca.onActivate = f.(func(*CompiledAura))
	}
	if aura.OnDeactivate != "" {
		f, err := bridge.Eval(aura.OnDeactivate)
		if err != nil {
			return nil, err
		}
		ca.onDeactivate = f.(func(*CompiledAura))
	}
	return ca, nil
}

func newCompiledAura(aura *Aura) *CompiledAura {
	return &CompiledAura{
		Enables:  make(map[string]func(string, []interface{}) bool),
		Disables: make(map[string]func(string, []interface{}) bool),
		events:   make(map[string]func(*CompiledAura, string, []interface{})),
	}
}
