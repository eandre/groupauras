package wow

import "github.com/eandre/lunar/lua"

func GetTime() float32 {
	return lua.Raw(`GetTime()`).(float32)
}
