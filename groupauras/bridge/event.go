package bridge

func RegisterEvent(event string, f func(string, []interface{})) {

}

func UnregisterEvent(event string, f func(string, []interface{})) {

}

func RegisterUpdate(f func(dt float32)) {

}

func UnregisterUpdate(f func(dt float32)) {

}
