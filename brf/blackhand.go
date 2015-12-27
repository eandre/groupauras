package brf

import "github.com/eandre/sbm"

type Blackhand struct {
	phase int
}

func (e *Blackhand) Start() {

}

func init() {
	sbm.RegisterEncounter("Blackhand", 1583, func() sbm.Encounter {
		return &Blackhand{phase: 1}
	})
}

func MyFunc() {

}
