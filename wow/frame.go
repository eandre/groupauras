package wow

func CreateFrame(typ string) *Frame {
	return &Frame{
		typ:    typ,
		parent: nil,
	}
}

type Frame struct {
	typ    string
	parent *Frame
}

func (f *Frame) SetParent(parent *Frame) {
	f.parent = parent
}

func (f *Frame) GetParent() *Frame {
	return f.parent
}
