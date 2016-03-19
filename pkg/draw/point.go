package draw

import "github.com/eandre/groupauras/shim/widget"

type Point struct {
	Pos   Position
	frame *pointFrame
}

func NewPoint(pos Position) *Point {
	return &Point{
		Pos:   pos,
		frame: getPointFrame(),
	}
}

type pointFrame struct {
	frame   widget.Frame
	texture widget.Texture
}

var pointFrameCache map[*pointFrame]bool

func getPointFrame() *pointFrame {
	for f := range pointFrameCache {
		delete(pointFrameCache, f)
		return f
	}

	f := &pointFrame{}
	f.frame = widget.CreateFrame(nil)
	f.frame.SetFrameStrata(widget.StrataMedium)
	f.texture = f.frame.CreateTexture()
	f.texture.SetAllPoints(nil)
	return f
}

func freePointFrame(f *pointFrame) {
	pointFrameCache[f] = true
}
