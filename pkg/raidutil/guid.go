package raidutil

import (
	"github.com/eandre/lunar-wow/pkg/luastrings"
	"github.com/eandre/lunar-wow/pkg/luatable"
	"github.com/eandre/lunar-wow/pkg/luautil"
	"github.com/eandre/lunar-wow/pkg/wow"
)

var guidToUnit = make(map[wow.GUID]wow.UnitID)

func GUIDToUnitID(guid wow.GUID) (unit wow.UnitID, ok bool) {
	uid := guidToUnit[guid]
	return uid, uid != ""
}

func guidToUnitUpdate(event string, args []interface{}) {
	luatable.Wipe(guidToUnit)

	update := func(unit wow.UnitID) {
		guid := wow.UnitGUID(unit)
		if !luautil.IsNil(guid) {
			guidToUnit[guid] = unit
		}
	}

	update("player")

	prefix := "party"
	inRaid := wow.IsInRaid()
	if inRaid {
		prefix = "raid"
	}
	num := wow.GetNumGroupMembers()
	for i := 1; i <= num; i++ {
		if i == num && !inRaid {
			// "partyN" is not a valid unit id, and we've already covered "player"
			return
		}
		update(wow.UnitID(prefix + luastrings.ToString(i)))
	}
}

func init() {
	wow.RegisterEvent("GROUP_ROSTER_UPDATE", guidToUnitUpdate)
	wow.RegisterEvent("RAID_ROSTER_UPDATE", guidToUnitUpdate)
	wow.RegisterEvent("PLAYER_ENTERING_WORLD", guidToUnitUpdate)
}
