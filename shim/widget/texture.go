package widget

type Texture interface {
	LayeredRegion

	GetTexture() string
	SetTexture(texture string)
}
