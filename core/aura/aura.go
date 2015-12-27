package aura
import "github.com/eandre/sbm/wow/persist"

type Aura struct {
	Encounter int
	Name string
}

func ForEncounter(id int) []*Aura {
	var result []*Aura
	for _, aura := range auras {
		if aura.Encounter == id {
			result = append(result, aura)
		}
	}
	return result
}

func Add(aura *Aura) {
	auras = append(auras, aura)
}

var auras []*Aura

func init() {
	restored, err := persist.Restore("auras")
	if err != nil {
		panic("Could not restore auras: " + err.Error())
	}

	if restored != nil {
		// We have auras; restore them
		auras = restored.([]*Aura)
	}
}
