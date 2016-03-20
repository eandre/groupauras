package draw

import "github.com/eandre/groupauras/shim/widget"

type PointCfg struct {
	Pos     Position
	Texture string
}

type Point struct {
	cfg   *PointCfg
	frame *pointFrame
}

func NewPoint(cfg *PointCfg) *Point {
	p := &Point{
		cfg:   cfg,
		frame: getPointFrame(),
	}
	p.frame.SetTexture(cfg.Texture)
	p.frame.frame.SetSize(20, 20)
	markPointActive(p)
	return p
}

func (p *Point) Free() {
	markPointInactive(p)
	freePointFrame(p.frame)
}

func (p *Point) update() {
	x, y, inst := p.cfg.Pos.Pos()
	dx, dy, show := displayOffset(x, y, inst)
	if !show {
		p.frame.frame.Hide()
		return
	}
	p.frame.frame.Show()

	// TODO(eandre) Determine if we should update?
	p.frame.SetPosition(dx, dy)
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

func (f *pointFrame) SetPosition(dx, dy float32) {
	f.frame.ClearAllPoints()
	f.frame.SetPoint("CENTER", f.frame.GetParent(), "CENTER", dx, dy)
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
