package activity

import "github.com/eandre/sbm/groupauras/aura"

type Tracker struct {
	activeAuras map[*aura.CompiledAura]bool
}

func (t *Tracker) Enable(a *aura.CompiledAura) {
	// See if it's already enabled
	if t.activeAuras[a] {
		return
	}
	t.activeAuras[a] = true
	a.Enable()
}

func (t *Tracker) Disable(a *aura.CompiledAura) {
	// Make sure it's enabled
	if !t.activeAuras[a] {
		return
	}
	a.Disable()
	delete(t.activeAuras, a)
}
