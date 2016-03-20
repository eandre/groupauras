package draw

import (
	"github.com/eandre/groupauras/shim/bridge"
	"github.com/eandre/groupauras/shim/hbd"
	"github.com/eandre/groupauras/shim/luamath"
	"github.com/eandre/groupauras/shim/widget"
)

var RotateMap = true

var canvas widget.Frame
var pixelsPerYard float32

func init() {
	uiparent := widget.UIParent()
	canvas = widget.CreateFrame(uiparent)
	canvas.SetSize(uiparent.GetWidth(), uiparent.GetHeight())
	canvas.SetPoint("CENTER", uiparent, "CENTER", 0, 0)
}

func displayOffset(x, y hbd.WorldCoord, inst hbd.InstanceID) (dx, dy float32, show bool) {
	px, py, i := hbd.PlayerWorldPosition()
	if i != inst {
		return 0, 0, false
	}

	dx, dy = float32(x-px), float32(y-py)

	if RotateMap {
		facing := bridge.PlayerFacing()
		fsin := luamath.Sin(facing)
		fcos := luamath.Cos(facing)
		tx, ty := dx, dy
		dx = tx*fcos - ty*fsin
		dy = tx*fsin + ty*fcos
	}

	return -dx * pixelsPerYard, dy * pixelsPerYard, true
}

func init() {
	pixelsPerYard = widget.UIParent().GetHeight() * 0.48 / 100
}
