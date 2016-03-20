package widget

type BlendMode string

const (
	BlendAdd      BlendMode = "ADD"
	BlendAlphaKey BlendMode = "ALPHAKEY"
	BlendBlend    BlendMode = "BLEND"
	BlendDisable  BlendMode = "DISABLE"
)

type Texture interface {
	LayeredRegion

	GetTexture() string
	SetTexture(texture string)
	SetTexCoord(coords ...float32)
	GetTexCoord() (ULx, ULy, LLx, LLy, URx, URy, LRx, LRy float32)

	GetBlendMode() BlendMode
	SetBlendMode(mode BlendMode)
}
