package widget

type ObjectType string

type UIObject interface {
	GetName() string
	GetObjectType() ObjectType
	IsObjectType(other ObjectType) bool
}

type ParentedObject interface {
	UIObject
	GetParent() UIObject
}

type ScriptType string

type ScriptObject interface {
	HasScript(typ ScriptType) bool
	GetScript(typ ScriptType) (handler interface{})
	SetScript(typ ScriptType, handler interface{})
	HookScript(typ ScriptType, handler interface{})
}
