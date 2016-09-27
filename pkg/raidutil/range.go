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
	if !wow.IsInRaid() {
		return
	}

	num := wow.GetNumGroupMembers()
	for i := 1; i <= num; i++ {
		src := wow.UnitID("raid" + luastrings.ToString(i))
		srcGUID := wow.UnitGUID(src)
		srcX, srcY, srcInst := hbd.UnitWorldPosition(src)
		var result []*RangeResult

		for j := 1; j < i; j++ {
			dst := wow.UnitID("raid" + luastrings.ToString(j))
			dstX, dstY, dstInst := hbd.UnitWorldPosition(dst)
			if dstInst == srcInst {
				yards := hbd.WorldDistance(srcInst, srcX, srcY, dstX, dstY)
				result = append(result, &RangeResult{
					UnitID: dst,
					GUID:   wow.UnitGUID(dst),
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
