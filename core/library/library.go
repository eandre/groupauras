package library

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/core/library/transfer"
)

type Listener interface {
	OnAuraChanged(aura *aura.Aura, enabled bool)
}

type Library struct {
	listeners []Listener
	auras     []*aura.Aura
	enabled   map[string]bool // aura id -> enabled
}

func New(auras []*aura.Aura) *Library {
	l := &Library{
		auras: auras,
	}
	transfer.RegisterListener(l)
	return l
}

func (l *Library) RegisterListener(ln Listener) {
	l.listeners = append(l.listeners, ln)
}

func (l *Library) OnAuraReceived(new *aura.Aura) {
	for i, old := range l.auras {
		if old.ID == new.ID {
			if old.Revision < new.Revision {
				l.auras[i] = new
			}
			return
		}
	}
	l.auras = append(l.auras, new)
}

func (l *Library) Auras() []*aura.Aura {
	return l.auras
}
