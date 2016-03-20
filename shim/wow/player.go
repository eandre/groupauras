package wow

import "github.com/eandre/lunar/lua"

func PlayerFacing() float32 {
	return lua.Raw(`GetPlayerFacing()`).(float32)
}
