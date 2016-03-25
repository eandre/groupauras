package acecfg

type Option interface {
	typ() string
	data() map[string]interface{}
}

type Info struct{}

// HACK to get static typing without creating methods for each attribute
type baseFieldsOption struct {
	Name     string
	Desc     string
	Validate func(info *Info, value interface{}) string
	Disabled func(info *Info, value interface{}) bool
}

func (o *baseFieldsOption) typ() string                  { return "" }
func (o *baseFieldsOption) data() map[string]interface{} { return nil }
func baseData(o *baseFieldsOption) map[string]interface{} {
	return map[string]interface{}{
		"type":     o.typ(),
		"name":     o.Name,
		"desc":     o.Desc,
		"validate": o.Validate,
		"disabled": o.Disabled,
	}
}

type Group struct {
	Name     string
	Desc     string
	Validate func(info *Info, value interface{}) string
	Disabled func(info *Info, value interface{}) bool

	Children map[string]Option
}

func (g *Group) typ() string { return "group" }
func (g *Group) data() map[string]interface{} {
	args := map[string]interface{}{}
	for key, opt := range g.Children {
		args[key] = serializeOption(opt)
	}
	return map[string]interface{}{
		"args": args,
	}
}

type Range struct {
	Name     string
	Desc     string
	Validate func(info *Info, value interface{}) string
	Disabled func(info *Info, value interface{}) bool

	Min       float32
	Max       float32
	SoftMin   float32
	SoftMax   float32
	Step      float32
	BigStep   float32
	IsPercent bool
	Get       func(info *Info) float32
	Set       func(info *Info, value float32)
}

func (r *Range) typ() string { return "range" }
func (r *Range) data() map[string]interface{} {
	return map[string]interface{}{
		"min":       r.Min,
		"max":       r.Max,
		"softMin":   r.SoftMin,
		"softMax":   r.SoftMax,
		"step":      r.Step,
		"bigStep":   r.BigStep,
		"get":       r.Get,
		"set":       r.Set,
		"isPercent": r.IsPercent,
	}
}

type Toggle struct {
	Name     string
	Desc     string
	Validate func(info *Info, value interface{}) string
	Disabled func(info *Info, value interface{}) bool

	Get func(info *Info) bool
	Set func(info *Info, value bool)
}

func (t *Toggle) typ() string { return "toggle" }
func (t *Toggle) data() map[string]interface{} {
	return map[string]interface{}{
		"get": t.Get,
		"set": t.Set,
	}
}

func serializeOption(opt Option) map[string]interface{} {
	o := opt.(*baseFieldsOption)
	data := baseData(o)
	for k, v := range opt.data() {
		data[k] = v
	}
	return data
}
