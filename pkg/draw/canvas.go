package draw

import "github.com/eandre/groupauras/shim/widget"

var canvas widget.Frame

func init() {
	uiparent := widget.UIParent()
	canvas = widget.CreateFrame(uiparent)
	canvas.SetSize(uiparent.GetWidth(), uiparent.GetHeight())
	canvas.SetPoint("CENTER", uiparent, "CENTER", 0, 0)
}
