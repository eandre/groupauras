package draw

import (
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/lunar-shim/hbd"
	"github.com/eandre/lunar-wow/pkg/luastrings"
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
	// Group unit ids are not stable over time; if we have
	// a group unit id, convert it to a guid first.
	if luastrings.HasPrefix(string(unitID), "raid") || luastrings.HasPrefix(string(unitID), "party") {
		guid := wow.UnitGUID(unitID)
		return &guidPosition{guid}
	}
	return &unitIDPosition{unitID}
}

func GUIDPosition(guid wow.GUID) Position {
	return &guidPosition{guid}
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

type unitIDPosition struct {
	unitID wow.UnitID
}

func (up *unitIDPosition) Static() bool {
	return false
}

func (up *unitIDPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	return hbd.UnitWorldPosition(up.unitID)
}

type guidPosition struct {
	guid wow.GUID
}

func (up *guidPosition) Static() bool {
	return false
}

func (up *guidPosition) Pos() (x, y hbd.WorldCoord, inst hbd.InstanceID) {
	uid, ok := raidutil.GUIDToUnitID(up.guid)
	if !ok {
		return 0, 0, 0
	}
	return hbd.UnitWorldPosition(uid)
}
