package widget

type LoopType string

const (
	LoopBounce LoopType = "BOUNCE"
	LoopNone   LoopType = "NONE"
	LoopRepeat LoopType = "REPEAT"
)

type LoopState string

const (
	LoopStateForward LoopState = "FORWARD"
	LoopStateNone    LoopState = "NONE"
	LoopStateReverse LoopState = "REPEAT"
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
	GetLoopState() LoopState

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

	Play()
	Pause()
	Stop()
	IsPlaying() bool
	IsPaused() bool
	IsStopped() bool

	SetDuration(duration float32)
	GetDuration() float32
	SetOrder(order int)
	GetOrder() int
	GetProgress() float32

	GetRegionParent() Region
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

type ScaleAnimation interface {
	Animation

	SetScale(xScale, yScale float32)
	GetScale() (xScale, yScale float32)

	SetOrigin(point AnchorPoint, xOff, yOff float32)
	GetOrigin() (point AnchorPoint, xOff, yOff float32)
}

type AlphaAnimation interface {
	Animation

	SetFromAlpha(from float32)
	SetToAlpha(to float32)
	SetChange(change float32)
	GetChange() float32
	GetFromAlpha() float32
	GetToAlpha() float32
}
