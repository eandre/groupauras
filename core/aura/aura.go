package aura

// Note that these types cannot have methods; due to how we store
// them in SavedVariables we cannot easily deserialize them to
// proper objects with metatables.

type Aura struct {
	ID       string
	Name     string
	Author   string
	Revision int

	Enables      map[string]string
	Disables     map[string]string
	Events       map[string]string
	OnUpdate     string
	OnActivate   string
	OnDeactivate string

	Raw string
	Sig string
}

type NewAura struct {
	ID       string
	Name     string
	Author   string
	Revision int

	Enables  map[string]string
	Disables map[string]string
}
