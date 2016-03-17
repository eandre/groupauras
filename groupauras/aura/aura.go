package aura

type Aura struct {
	Name     string
	Enables  map[string]string
	Disables map[string]string
	Events   map[string]string
	Update   string
}

func New(name string) *Aura {
	return &Aura{
		Name:     name,
		Enables:  make(map[string]string),
		Disables: make(map[string]string),
		Events:   make(map[string]string),
	}
}
