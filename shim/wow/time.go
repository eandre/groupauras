package wow

import "github.com/eandre/lunar/lua"

func GetTime() (secs float32) {
	return lua.Raw(`GetTime()`).(float32)
}
