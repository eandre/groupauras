package acecomm

type StatusFunc func(key interface{}, currBytes, totalBytes int)

func SendCommMessage(prefix, msg, channel string, target interface{}, prio string, f StatusFunc, key interface{}) {

}

type Handler func(prefix, message, distribution, sender string)

func RegisterComm(prefix string, handler Handler) {

}
