package activity

import "github.com/eandre/groupauras/aura"

type Tracker struct {
	activeAuras map[*aura.CompiledAura]bool
}

func (t *Tracker) Activate(a *aura.CompiledAura) {
	// See if it's already active
	if t.activeAuras[a] {
		return
	}
	t.activeAuras[a] = true
	a.Activate()
}

func (t *Tracker) Deactivate(a *aura.CompiledAura) {
	// Make sure it's active
	if !t.activeAuras[a] {
		return
	}
	a.Deactivate()
	delete(t.activeAuras, a)
}
