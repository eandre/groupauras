package bridge

type EventListener func(event string, args []interface{})

func RegisterEvent(event string, listener EventListener) {

}

func UnregisterEvent(event string, listener EventListener) {

}
