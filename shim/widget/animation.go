package widget

type LoopType string

const (
	LoopBounce LoopType = "BOUNCE"
	LoopNone   LoopType = "NONE"
	LoopRepeat LoopType = "REPEAT"
)

type AnimationGroup interface {
	ScriptObject
	ParentedObject

	Finish()
	Stop()
	Play()
	Pause()
	IsPlaying() bool
	IsPaused() bool
	IsDone() bool
	IsPendingFinish() bool

	SetLooping(loopType LoopType)
	GetLooping() LoopType

	CreateAnimation(typ AnimationType) Animation
}

type AnimationType string

const (
	AnimationRotation    AnimationType = "rotation"
	AnimationScale       AnimationType = "scale"
	AnimationTranslation AnimationType = "translation"
	AnimationAlpha       AnimationType = "alpha"
	AnimationPath        AnimationType = "path"
)

type Animation interface {
	ParentedObject
	ScriptObject

	SetDuration(duration float32)
	GetDuration() float32
}

type RotationAnimation interface {
	Animation

	SetDegrees(degrees float32)
	GetDegrees() float32
	SetRadians(radians float32)
	GetRadians() float32

	SetOrigin(point AnchorPoint, xOff, yOff float32)
	GetOrigin() (point AnchorPoint, xOff, yOff float32)
}
