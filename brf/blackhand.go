package brf

import "github.com/eandre/sbm"
import "github.com/eandre/sbm/wow"

type Blackhand struct {
	phase int
}

func (e *Blackhand) Start() {

}

func init() {
	wow.TestWow(5)
	sbm.RegisterEncounter("Blackhand", 1583, func() sbm.Encounter {
		return &Blackhand{phase: 1}
	})
}

func MyFunc() {

}
