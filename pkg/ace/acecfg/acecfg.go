package acecfg

type Option interface {
	option()
}

type Info struct{}

type Group struct {
	Name  string `luaname:"name"`
	Desc  string `luaname:"desc"`
	Order int    `luaname:"order"`
	Type  string `luaname:"type" luadefault:"\"group\""`

	Validate func(info *Info, value interface{}) string `luaname:"validate'`
	Disabled func(info *Info, value interface{}) bool   `luaname:"disabled"`

	Children map[string]Option `luaname:"args"`
}

func (g *Group) option() {}

type Range struct {
	Name  string `luaname:"name"`
	Desc  string `luaname:"desc"`
	Order int    `luaname:"order"`
	Type  string `luaname:"type" luadefault:"\"range\""`

	Validate func(info *Info, value interface{}) string `luaname:"validate'`
	Disabled func(info *Info, value interface{}) bool   `luaname:"disabled"`

	Min       float32 `luaname:"min"`
	Max       float32 `luaname:"max"`
	SoftMin   float32 `luaname:"softMin"`
	SoftMax   float32 `luaname:"softMax"`
	Step      float32 `luaname:"step"`
	BigStep   float32 `luaname:"bigStep"`
	IsPercent bool    `luaname:"isPercent"`

	Get func(info *Info) float32        `luaname:"min"`
	Set func(info *Info, value float32) `luaname:"min"`
}

func (r *Range) option() {}

type Toggle struct {
	Name  string `luaname:"name"`
	Desc  string `luaname:"desc"`
	Order int    `luaname:"order"`
	Type  string `luaname:"type" luadefault:"\"toggle\""`

	Validate func(info *Info, value interface{}) string `luaname:"validate'`
	Disabled func(info *Info, value interface{}) bool   `luaname:"disabled"`

	Get func(info *Info) bool        `luaname:"get"`
	Set func(info *Info, value bool) `luaname:"set"`
}

func (t *Toggle) option() {}

type Button struct {
	Name  string `luaname:"name"`
	Desc  string `luaname:"desc"`
	Order int    `luaname:"order"`
	Type  string `luaname:"type" luadefault:"\"execute\""`

	Validate func(info *Info, value interface{}) string `luaname:"validate'`
	Disabled func(info *Info, value interface{}) bool   `luaname:"disabled"`

	Func func(info *Info) `luaname:"func"`
}

func (b *Button) option() {}
