package classutil

import (
	"github.com/eandre/lunar-wow/pkg/wow"
	"github.com/eandre/lunar/lua"
)

type raidClassColor struct {
	r        float32
	g        float32
	b        float32
	colorStr string
}

func Color(unit wow.UnitID) (r, g, b float32) {
	_, file := wow.UnitClassBase(unit)

	_ = file

	tbl := lua.Raw(`RAID_CLASS_COLORS[file]`).(*raidClassColor)
	return tbl.r, tbl.g, tbl.b
}
