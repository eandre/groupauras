package luamath

import "github.com/eandre/lunar/lua"

var floor = lua.Raw(`math.floor`).(func(val float32) float32)

func Floor(val float32) float32 {
	return floor(val)
}
