package widget

import "github.com/eandre/lunar/lua"

type FrameStrata string

const (
	StrataBackground       FrameStrata = "BACKGROUND"
	StrataLow              FrameStrata = "LOW"
	StrataMedium           FrameStrata = "MEDIUM"
	StrataHigh             FrameStrata = "HIGH"
	StrataDialog           FrameStrata = "DIALOG"
	StrataFullscreen       FrameStrata = "FULLSCREEN"
	StrataFullscreenDialog FrameStrata = "FULLSCREEN_DIALOG"
	StrataTooltip          FrameStrata = "TOOLTIP"
)

type Frame interface {
	VisibleRegion
	ScriptObject

	CreateTexture() Texture

	GetFrameStrata() FrameStrata
	SetFrameStrata(strata FrameStrata)
}

func CreateFrame(parent UIObject) Frame {
	return lua.Raw(`CreateFrame("Frame", nil, parent)`).(Frame)
}

func UIParent() Frame {
	return lua.Raw("UIParent").(Frame)
}
