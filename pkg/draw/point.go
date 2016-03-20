package draw

import "github.com/eandre/groupauras/shim/widget"

type Point struct {
	Pos   Position
	frame *pointFrame
}

func NewPoint(pos Position, texture string) *Point {
	point := &Point{
		Pos:   pos,
		frame: getPointFrame(),
	}
	point.frame.SetTexture(texture)
	point.frame.frame.SetSize(20, 20)
	return point
}

func (p *Point) Free() {
	freePointFrame(p.frame)
}

type pointFrame struct {
	frame   widget.Frame
	texture widget.Texture
}

func (f *pointFrame) SetTexture(texture string) {
	entry := textureMap[texture]
	if entry != nil {
		f.texture.SetTexture(entry.Texture)
		f.texture.SetTexCoord(entry.TexCoords...)
		f.texture.SetBlendMode(entry.Blend)
	} else {
		f.texture.SetTexture(texture)
		f.texture.SetTexCoord(0, 1, 0, 1)
		f.texture.SetBlendMode(widget.BlendBlend)
	}
}

var pointFrameCache map[*pointFrame]bool

func getPointFrame() *pointFrame {
	for f := range pointFrameCache {
		delete(pointFrameCache, f)
		return f
	}
	return newPointFrame()
}

func freePointFrame(f *pointFrame) {
	pointFrameCache[f] = true
}

func newPointFrame() *pointFrame {
	f := &pointFrame{}
	f.frame = widget.CreateFrame(canvas)
	f.frame.SetFrameStrata(widget.StrataMedium)
	f.texture = f.frame.CreateTexture()
	f.texture.SetAllPoints(f.frame)
	f.texture.SetDrawLayer(widget.LayerArtwork, 0)
	return f
}
