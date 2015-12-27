package encounter
import (
	"github.com/eandre/sbm/core/aura"
	"github.com/eandre/sbm/wow/event"
)

type Encounter struct {
	ID int
	Phase int
	Auras []*aura.Aura
}

var activeEncounters []*Encounter

func onEncounterStart(event string, args ...interface{}) {
	name := args[0].(string)
	id := args[1].(int)
	auras := aura.ForEncounter(id)

	if len(auras) != 0 {
		println("Started encounter", name)
		activeEncounters = append(activeEncounters, &Encounter{
			ID: id,
			Phase: 1,
			Auras: auras,
		})
	}
}

func init() {
	event.Register(event.EventEncounterStart, onEncounterStart)
}