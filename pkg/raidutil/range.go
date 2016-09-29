package raidutil

import (
	"github.com/eandre/lunar-shim/hbd"
	"github.com/eandre/lunar-wow/pkg/luastrings"
	"github.com/eandre/lunar-wow/pkg/luatable"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type RangeResult struct {
	UnitID wow.UnitID
	GUID   wow.GUID
	Yards  float32
}

var ranges = make(map[wow.GUID][]*RangeResult)

func GroupMembersWithin(guid wow.GUID, yards float32) []*RangeResult {
	var result []*RangeResult
	r := ranges[guid]
	if len(r) == 0 {
		return nil
	}

	for _, res := range r {
		if res.Yards <= yards {
			result = append(result, res)
		}
	}
	return result
}

func updateRanges() {
	luatable.Wipe(ranges)
	inRaid := wow.IsInRaid()
	prefix := "party"
	if inRaid {
		prefix = "raid"
	}

	num := wow.GetNumGroupMembers()
	for i := 1; i <= num; i++ {
		src := wow.UnitID(prefix + luastrings.ToString(i))
		if !inRaid && i == num {
			src = "player"
		}
		srcGUID := wow.UnitGUID(src)
		srcX, srcY, srcInst := hbd.UnitWorldPosition(src)
		var result []*RangeResult
		for j := 1; j < num; j++ {
			dst := wow.UnitID(prefix + luastrings.ToString(j))
			dstX, dstY, dstInst := hbd.UnitWorldPosition(dst)
			if dstInst == srcInst {
				yards := hbd.WorldDistance(srcInst, srcX, srcY, dstX, dstY)
				dstGUID := wow.UnitGUID(dst)
				result = append(result, &RangeResult{
					UnitID: dst,
					GUID:   dstGUID,
					Yards:  yards,
				})
				ranges[dstGUID] = append(ranges[dstGUID], &RangeResult{
					UnitID: src,
					GUID:   srcGUID,
					Yards:  yards,
				})
			}
		}

		ranges[srcGUID] = result
	}
}

var elapsed float32

func onUpdate(dt float32) {
	elapsed += dt
	if elapsed < 0.1 {
		return
	}
	elapsed -= 0.1
	updateRanges()
}

func init() {
	wow.RegisterUpdate(onUpdate)
}
