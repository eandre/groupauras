package widget

type AnchorPoint string

type Region interface {
	ParentedObject

	GetTop() float32
	GetLeft() float32
	GetRight() float32
	GetBottom() float32
	GetRect() (left, bottom, width, height float32)

	GetWidth() float32
	GetHeight() float32
	GetSize() (width, height float32)
	SetWidth(float32)
	SetHeight(float32)
	SetSize(width, height float32)

	GetNumPoints() int
	GetPoint(idx int) (point AnchorPoint, relativeTo Region, relativePoint AnchorPoint, xOffset, yOffset float32)
	SetPoint(point AnchorPoint, relativeTo Region, relativePoint AnchorPoint, xOffset, yOffset float32)
	SetAllPoints(region Region)
	ClearAllPoints()

	IsDragging() bool
	IsMouseOver(topOff, leftOff, bottomOff, rightOff float32) bool

	CreateAnimationGroup() *AnimationGroup
	// GetAnimationGroups() []*AnimationGroup -- returns "..."
	StopAnimating()

	CanChangeProtectedState() bool
}

type VisibleRegion interface {
	Region

	Show()
	Hide()
	IsShown() bool
	IsVisible() bool

	GetAlpha() float32
	SetAlpha(float32)
}

type DrawLayer string

const (
	LayerArtwork    DrawLayer = "ARTWORK"
	LayerBackground DrawLayer = "BACKGROUND"
	LayerBorder     DrawLayer = "BORDER"
	LayerHighlight  DrawLayer = "HIGHLIGHT"
	LayerOverlay    DrawLayer = "OVERLAY"
)

type LayeredRegion interface {
	VisibleRegion

	GetDrawLayer() (layer DrawLayer, sublayer int)
	SetDrawLayer(layer DrawLayer, sublayer int)
	SetVertexColor(r, g, b, a float32)
}
