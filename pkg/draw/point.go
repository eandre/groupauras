package draw

import (
	"github.com/eandre/groupauras/shim/widget"
	"github.com/eandre/groupauras/shim/wow"
)

type PointCfg struct {
	Pos     Position
	Texture string

	// One of SizeYards and SizePixels must be set.
	// SizeYards takes precedence if both are set.
	SizeYards  float32
	SizePixels int

	// If set, the texture will rotate
	RotateDegrees float32
	RotateSpeed   float32

	// Number of seconds the point will exist for.
	// If zero, it will never expire.
	Duration float32
}

type Point struct {
	cfg   *PointCfg
	frame *pointFrame

	deadline float32
}

func NewPoint(cfg *PointCfg) *Point {
	p := &Point{
		cfg:   cfg,
		frame: getPointFrame(),
	}
	p.SetTexture(cfg.Texture)
	p.SetSize(cfg.SizeYards, cfg.SizePixels)

	if cfg.RotateSpeed != 0 {
		p.Rotate(cfg.RotateDegrees, cfg.RotateSpeed)
	}
	if cfg.Duration != 0 {
		p.SetDuration(cfg.Duration)
	}

	markPointActive(p)
	return p
}

func (p *Point) Free() {
	p.frame.Free()
	markPointInactive(p)
	freePointFrame(p.frame)
}

func (p *Point) Rotate(degrees, speed float32) {
	p.frame.Rotate(degrees, speed)
}

func (p *Point) SetTexture(texture string) {
	p.frame.SetTexture(texture)
}

func (p *Point) SetSize(sizeYards float32, sizePixels int) {
	if sizeYards != 0 {
		// Multiply by two since it's a radius
		p.frame.SetSize(sizeYards * pixelsPerYard() * 2)
	} else {
		p.frame.SetSize(float32(sizePixels))
	}
}

func (p *Point) SetDuration(secs float32) {
	p.deadline = wow.GetTime() + secs
}

func (p *Point) update() {
	if p.deadline != 0 && wow.GetTime() > p.deadline {
		p.Free()
		return
	}

	x, y, inst := p.cfg.Pos.Pos()
	dx, dy, show := displayOffset(x, y, inst)
	if !show {
		p.Free()
		return
	}
	p.frame.frame.Show()

	// TODO(eandre) Determine if we should update?
	p.frame.SetPosition(dx, dy)
}

type pointFrame struct {
	frame            widget.Frame
	texture          widget.Texture
	repeatAnimations widget.AnimationGroup
	rotate           widget.RotationAnimation

	texDef *textureDef // may be nil
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
	f.texDef = entry
}

func (f *pointFrame) SetSize(pixels float32) {
	if f.texDef != nil {
		pixels *= f.texDef.SizeScalar
	}
	f.frame.SetSize(pixels, pixels)
}

func (f *pointFrame) SetPosition(dx, dy float32) {
	f.frame.ClearAllPoints()
	f.frame.SetPoint("CENTER", f.frame.GetParent(), "CENTER", dx, dy)
}

func (f *pointFrame) Rotate(degrees, speed float32) {
	norm := 360 / degrees
	speed = speed * norm
	degrees = -360
	if speed < 0 {
		speed = speed * -1
		degrees = 360
	}

	f.rotate.SetDuration(speed)
	f.rotate.SetDegrees(degrees)
	f.repeatAnimations.Play()
}

func (f *pointFrame) Reset() {
	f.frame.StopAnimating()
	f.frame.Show()
	f.frame.SetAlpha(1)
	f.repeatAnimations.Stop()
}

func (f *pointFrame) Free() {
	f.frame.Hide()
	f.frame.StopAnimating()
}

func newPointFrame() *pointFrame {
	f := &pointFrame{}
	f.frame = widget.CreateFrame(canvas)
	f.frame.SetFrameStrata(widget.StrataMedium)
	f.texture = f.frame.CreateTexture()
	f.texture.SetAllPoints(f.frame)
	f.texture.SetDrawLayer(widget.LayerArtwork, 0)

	f.repeatAnimations = f.frame.CreateAnimationGroup()
	f.repeatAnimations.SetLooping(widget.LoopRepeat)
	f.rotate = f.repeatAnimations.CreateAnimation(widget.AnimationRotation).(widget.RotationAnimation)
	return f
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
