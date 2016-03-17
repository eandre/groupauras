package aura

import "github.com/eandre/sbm/groupauras/bridge"

type CompiledAura struct {
	Aura     *Aura
	Enables  map[string]func(event string, args []interface{}) bool
	Disables map[string]func(event string, args []interface{}) bool
	Events   map[string]func(event string, args []interface{})
	Update   func(dt float32)
}

func Compile(aura *Aura) (*CompiledAura, error) {
	ca := newCompiledAura(aura)
	for event, src := range aura.Enables {
		f, err := bridge.EvalEventBool(src)
		if err != nil {
			return nil, err
		}
		ca.Enables[event] = f
	}
	for event, src := range aura.Disables {
		f, err := bridge.EvalEventBool(src)
		if err != nil {
			return nil, err
		}
		ca.Disables[event] = f
	}
	for event, src := range aura.Events {
		f, err := bridge.EvalEvent(src)
		if err != nil {
			return nil, err
		}
		ca.Events[event] = f
	}
	if aura.Update != "" {
		f, err := bridge.EvalUpdate(aura.Update)
		if err != nil {
			return nil, err
		}
		ca.Update = f
	}
	return ca, nil
}

func newCompiledAura(aura *Aura) *CompiledAura {
	return &CompiledAura{
		Enables:  make(map[string]func(event string, args []interface{}) bool),
		Disables: make(map[string]func(event string, args []interface{}) bool),
		Events:   make(map[string]func(event string, args []interface{})),
	}
}
