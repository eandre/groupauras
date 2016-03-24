package draw

import (
	"github.com/eandre/groupauras/shim/hbd"
	"github.com/eandre/groupauras/shim/luamath"
	"github.com/eandre/groupauras/shim/widget"
	"github.com/eandre/groupauras/shim/wow"
)

var RotateMap = true
var ZoomScale float32 = 50

var canvas widget.Frame
var maxSize float32

func init() {
	uiparent := widget.UIParent()
	canvas = widget.CreateFrame(uiparent)
	canvas.SetSize(uiparent.GetWidth(), uiparent.GetHeight())
	canvas.SetPoint("CENTER", uiparent, "CENTER", 0, 0)
	maxSize = widget.UIParent().GetHeight() * 0.48
}

func displayOffset(x, y hbd.WorldCoord, inst hbd.InstanceID) (dx, dy float32, show bool) {
	px, py, i := hbd.PlayerWorldPosition()
	if i != inst {
		return 0, 0, false
	}

	dx, dy = float32(x-px), float32(y-py)

	if RotateMap {
		facing := wow.PlayerFacing()
		fsin := luamath.Sin(facing)
		fcos := luamath.Cos(facing)
		tx, ty := dx, dy
		dx = tx*fcos - ty*fsin
		dy = tx*fsin + ty*fcos
	}

	ppy := pixelsPerYard()
	return -dx * ppy, dy * ppy, true
}

func pixelsPerYard() float32 {
	return maxSize / ZoomScale
}
