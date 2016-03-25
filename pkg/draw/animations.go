package draw

import "github.com/eandre/groupauras/shim/widget"

type animationState struct {
	regionParent widget.VisibleRegion
}

func (s *animationState) OnLoad(anim widget.Animation) {
	s.regionParent = anim.GetRegionParent().(widget.VisibleRegion)
}

func (s *animationState) Alpha(anim widget.Animation) {
	s.regionParent.SetAlpha(anim.GetProgress())
}

func (s *animationState) FullOpacity(anim widget.Animation) {
	s.regionParent.SetAlpha(1)
}

func (s *animationState) HideParent(anim widget.Animation) {
	anim.GetRegionParent().(widget.VisibleRegion).Hide()
}

func (s *animationState) Replay(anim widget.Animation) {
	anim.Play()
}
