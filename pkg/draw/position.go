package draw

import "github.com/eandre/groupauras/shim/hbd"

type Position interface {
	Static() bool
	Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID)
}

type staticPosition struct {
	x, y hbd.WorldCoord
	inst hbd.InstanceID
}

func (sp *staticPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	return sp.x, sp.y, sp.inst
}

func (sp *staticPosition) Static() bool {
	return true
}

type playerPosition struct{}

func (pp *playerPosition) Static() bool {
	return false
}

func (pp *playerPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	return hbd.PlayerWorldPosition()
}

type unitPosition struct {
	unitID string
}

func (up *unitPosition) Static() bool {
	return false
}

func (up *unitPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	return hbd.UnitWorldPosition(up.unitID)
}
