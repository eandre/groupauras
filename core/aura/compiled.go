package aura

import "github.com/eandre/lunar-wow/pkg/luautil"

type CompiledAura struct {
	Aura     *Aura
	Enables  map[string]func(ca *CompiledAura, event string, args []interface{}) bool
	Disables map[string]func(ca *CompiledAura, event string, args []interface{}) bool

	// Relevant for activate auras
	Events       map[string]func(ca *CompiledAura, event string, args []interface{})
	OnUpdate     func(ca *CompiledAura, dt float32)
	OnActivate   func(ca *CompiledAura)
	OnDeactivate func(ca *CompiledAura)
}

func Compile(rev *Aura) (*CompiledAura, error) {
	ca := newCompiledAura(rev)
	for event, src := range rev.Enables {
		f, err := luautil.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.Enables[event] = f.(func(*CompiledAura, string, []interface{}) bool)
	}
	for event, src := range rev.Disables {
		f, err := luautil.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.Disables[event] = f.(func(*CompiledAura, string, []interface{}) bool)
	}
	for event, src := range rev.Events {
		f, err := luautil.Eval(src)
		if err != nil {
			return nil, err
		}
		ca.Events[event] = f.(func(*CompiledAura, string, []interface{}))
	}
	if rev.OnUpdate != "" {
		f, err := luautil.Eval(rev.OnUpdate)
		if err != nil {
			return nil, err
		}
		ca.OnUpdate = f.(func(*CompiledAura, float32))
	}
	if rev.OnActivate != "" {
		f, err := luautil.Eval(rev.OnActivate)
		if err != nil {
			return nil, err
		}
		ca.OnActivate = f.(func(*CompiledAura))
	}
	if rev.OnDeactivate != "" {
		f, err := luautil.Eval(rev.OnDeactivate)
		if err != nil {
			return nil, err
		}
		ca.OnDeactivate = f.(func(*CompiledAura))
	}
	return ca, nil
}

func newCompiledAura(rev *Aura) *CompiledAura {
	return &CompiledAura{
		Aura:     rev,
		Enables:  make(map[string]func(*CompiledAura, string, []interface{}) bool),
		Disables: make(map[string]func(*CompiledAura, string, []interface{}) bool),
		Events:   make(map[string]func(*CompiledAura, string, []interface{})),
	}
}
