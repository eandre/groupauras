package draw

import (
	"github.com/eandre/lunar-shim/hbd"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type Position interface {
	Static() bool
	Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID)
}

func PlayerPosition() Position {
	return &playerPosition{}
}

func UnitPosition(unitID wow.UnitID) Position {
	return &unitPosition{unitID}
}

func StaticPlayerPosition() Position {
	x, y, inst := hbd.PlayerWorldPosition()
	return &staticPosition{x, y, inst}
}

func StaticUnitPosition(unitID wow.UnitID) Position {
	x, y, inst := hbd.UnitWorldPosition(unitID)
	return &staticPosition{x, y, inst}
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
	unitID wow.UnitID
}

func (up *unitPosition) Static() bool {
	return false
}

func (up *unitPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	return hbd.UnitWorldPosition(up.unitID)
}
