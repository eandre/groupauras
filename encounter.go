package sbm

type Encounter interface {
	Start()
}

func RegisterEncounter(name string, id int, factory func() Encounter) {

}
