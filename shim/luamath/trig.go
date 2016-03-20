package luamath

import "github.com/eandre/lunar/lua"

var cos = lua.Raw(`math.cos`).(func(val float32) float32)
var sin = lua.Raw(`math.sin`).(func(val float32) float32)

func Cos(angle float32) float32 {
	return cos(angle)
}

func Sin(angle float32) float32 {
	return sin(angle)
}
