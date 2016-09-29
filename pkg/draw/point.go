package draw

import (
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/lunar-wow/pkg/widget"
)

type PointCfg struct {
	// If non-nil, when ctx is cancelled, the point will be freed
	Ctx         context.Ctx
	Pos         Position
	Texture     string
	Text        string
	VertexColor []float32

	// One of SizeYards and SizePixels must be set.
	// SizeYards takes precedence if both are set.
	SizeYards  float32
	SizePixels int

	// If set, the texture will rotate
	RotateDegrees float32
	RotateSpeed   float32

	DrawLayer widget.DrawLayer
}

type Point struct {
	ctx   context.Ctx
	cfg   *PointCfg
	frame *pointFrame
}

func NewPoint(cfg *PointCfg) *Point {
	ctx := cfg.Ctx
	if ctx == nil {
		ctx = context.Base
	}
	p := &Point{
		ctx:   ctx,
		cfg:   cfg,
		frame: getPointFrame(),
	}
	p.SetTexture(cfg.Texture, cfg.DrawLayer)
	p.SetSize(cfg.SizeYards, cfg.SizePixels)
	p.SetText(cfg.Text)

	if c := cfg.VertexColor; len(c) != 0 {
		p.SetVertexColor(c[0], c[1], c[2], c[3])
	}

	if cfg.RotateSpeed != 0 {
		p.Rotate(cfg.RotateDegrees, cfg.RotateSpeed)
	}
	p.FadeIn()

	markPointActive(p)
	return p
}

func (p *Point) Free(skipAnimations bool) {
	p.frame.Free(skipAnimations)
	markPointInactive(p)
	freePointFrame(p.frame)
}

func (p *Point) Rotate(degrees, speed float32) {
	p.frame.Rotate(degrees, speed)
}

func (p *Point) SetTexture(texture string, layer widget.DrawLayer) {
	if layer == "" {
		layer = widget.LayerArtwork
	}
	p.frame.SetTexture(texture, layer)
}

func (p *Point) SetText(text string) {
	p.frame.SetText(text)
}

func (p *Point) SetVertexColor(r, g, b, a float32) {
	p.frame.texture.SetVertexColor(r, g, b, a)
}

func (p *Point) SetSize(sizeYards float32, sizePixels int) {
	if sizeYards != 0 {
		// Multiply by two since it's a radius
		p.frame.SetSize(sizeYards * pixelsPerYard() * 2)
	} else {
		p.frame.SetSize(float32(sizePixels))
	}
}

func (p *Point) FadeIn() {
	p.frame.FadeIn()
}

func (p *Point) update() {
	if p.ctx.Cancelled() {
		p.Free(false)
		return
	}

	x, y, inst := p.cfg.Pos.Pos()
	dx, dy, show := displayOffset(x, y, inst)
	if !show {
		p.Free(false) // should we just hide instead?
		return
	}

	// TODO(eandre) Determine if we should update?
	p.frame.SetPosition(dx, dy)
}

type pointFrame struct {
	frame widget.Frame

	text widget.FontString

	texture          widget.Texture
	repeatAnimations widget.AnimationGroup
	rotate           widget.RotationAnimation

	fadeInAnimations widget.AnimationGroup
	scaleOut         widget.ScaleAnimation
	fadeIn           widget.AlphaAnimation
	scaleIn          widget.ScaleAnimation

	fadeOutAnimations widget.AnimationGroup
	fadeOut           widget.AlphaAnimation

	texDef *textureDef // may be nil
}

func (f *pointFrame) SetTexture(texture string, layer widget.DrawLayer) {
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
	f.texture.SetDrawLayer(layer, 0)
	f.texDef = entry
}

func (f *pointFrame) SetText(text string) {
	f.text.SetText(text)
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
	f.texture.SetVertexColor(1, 1, 1, 1)
	f.text.Show()
}

func (f *pointFrame) Free(skipAnimations bool) {
	if skipAnimations {
		f.frame.Hide()
		f.frame.StopAnimating()
		f.text.Hide()
	} else {
		f.FadeOut()
	}
}

func (f *pointFrame) FadeIn() {
	f.frame.Show()
	f.fadeInAnimations.Play()
}

func (f *pointFrame) FadeOut() {
	f.fadeOutAnimations.Play()
	f.text.Hide()
}

func newPointFrame() *pointFrame {
	f := &pointFrame{}
	f.frame = widget.CreateFrame(canvas)
	f.frame.SetFrameStrata(widget.StrataMedium)

	f.text = widget.UIParent().CreateFontString()
	f.text.SetFontObject("GameFontNormal")
	f.text.SetPoint("CENTER", f.frame, "CENTER", 0, 10)
	f.text.SetTextColor(1, 1, 1, 1)

	f.texture = f.frame.CreateTexture()
	f.texture.SetAllPoints(f.frame)
	f.texture.SetDrawLayer(widget.LayerArtwork, 0)

	f.repeatAnimations = f.frame.CreateAnimationGroup()
	f.repeatAnimations.SetLooping(widget.LoopRepeat)
	f.rotate = f.repeatAnimations.CreateAnimation(widget.AnimationRotation).(widget.RotationAnimation)

	// Fade in animations
	f.fadeInAnimations = f.frame.CreateAnimationGroup()
	f.scaleOut = f.fadeInAnimations.CreateAnimation(widget.AnimationScale).(widget.ScaleAnimation)
	f.scaleOut.SetScale(1.5, 1.5)
	f.scaleOut.SetOrder(1)

	fadeInState := &animationState{}
	f.fadeIn = f.fadeInAnimations.CreateAnimation(widget.AnimationAlpha).(widget.AlphaAnimation)
	f.fadeIn.SetDuration(0.35)
	f.fadeIn.SetScript("OnPlay", func(anim widget.AlphaAnimation) {
		fadeInState.OnLoad(anim)
		f.fadeOutAnimations.Stop()
	})
	f.fadeIn.SetFromAlpha(0)
	f.fadeIn.SetToAlpha(1)
	f.fadeIn.SetScript("OnUpdate", fadeInState.Alpha)
	f.fadeIn.SetScript("OnStop", fadeInState.FullOpacity)
	f.fadeIn.SetOrder(2)

	f.scaleIn = f.fadeInAnimations.CreateAnimation(widget.AnimationScale).(widget.ScaleAnimation)
	f.scaleIn.SetDuration(0.35)
	f.scaleIn.SetScale(1/1.5, 1/1.5)
	f.scaleIn.SetOrder(2)

	// Fade out animations
	fadeOutState := &animationState{}
	f.fadeOutAnimations = f.frame.CreateAnimationGroup()
	f.fadeOut = f.fadeOutAnimations.CreateAnimation(widget.AnimationAlpha).(widget.AlphaAnimation)
	f.fadeOut.SetFromAlpha(1)
	f.fadeOut.SetToAlpha(0)
	f.fadeOut.SetDuration(0.25)
	f.fadeOut.SetScript("OnFinished", fadeOutState.HideParent)
	f.fadeOutAnimations.SetScript("OnPlay", func() {
		f.fadeInAnimations.Stop()
	})

	return f
}

var pointFrameCache = make(map[*pointFrame]bool)

func getPointFrame() *pointFrame {
	for f := range pointFrameCache {
		delete(pointFrameCache, f)
		f.Reset()
		return f
	}
	return newPointFrame()
}

func freePointFrame(f *pointFrame) {
	pointFrameCache[f] = true
}
