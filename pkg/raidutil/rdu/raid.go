package rdu

import (
	"github.com/eandre/groupauras/pkg/classutil"
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/lunar-wow/pkg/luastrings"
	"github.com/eandre/lunar-wow/pkg/widget"
	"github.com/eandre/lunar-wow/pkg/wow"
)

func ShowGroupMembers(ctx context.Ctx) {
	doShowGroupMembers()

	ctx.OnCancel(nil, func(_ context.Ctx, _ interface{}) {
		doHideGroupMembers()
	})
}

type drawingMember struct {
	Point   *draw.Point
	GUID    wow.GUID
	Counter int
}

var currMembers = make(map[wow.GUID]*drawingMember)
var memberCount = 0 // prevent quadratic complexity

var timer float32 = 0

func updateTextures(dt float32) {
	timer += dt
	if timer < 0.1 {
		return
	}
	timer -= 0.1

	for _, dm := range currMembers {
		unitID, ok := raidutil.GUIDToUnitID(dm.GUID)
		if ok {
			idx := wow.GetRaidTargetIndex(unitID)
			if idx >= 1 && idx <= 8 {
				// Marked; use a marker texture
				dm.Point.SetTexture("marker_"+luastrings.ToString(idx), widget.LayerOverlay)
				dm.Point.SetVertexColor(1, 1, 1, 1)
			} else {
				// Not marked; use party dot
				r, g, b := classutil.Color(unitID)
				dm.Point.SetTexture("party", widget.LayerOverlay)
				dm.Point.SetVertexColor(r, g, b, 1)
			}
		}
	}
}

func addMember(guid wow.GUID) {
	if dm := currMembers[guid]; dm != nil {
		dm.Counter += 1
		return
	}

	pt := draw.NewPoint(&draw.PointCfg{
		Pos:         draw.GUIDPosition(guid),
		Texture:     "party",
		VertexColor: []float32{1, 1, 1, 1}, // filled in by update
		SizeYards:   1,
		DrawLayer:   widget.LayerOverlay,
	})
	currMembers[guid] = &drawingMember{
		Point:   pt,
		GUID:    guid,
		Counter: 1,
	}
	memberCount += 1

	if memberCount == 1 {
		wow.RegisterUpdate(updateTextures)
	}
}

func removeMember(guid wow.GUID) {
	if dm := currMembers[guid]; dm != nil {
		dm.Counter -= 1
		if dm.Counter <= 0 {
			dm.Point.Free(false)
			delete(currMembers, guid)
			memberCount -= 1
		}
	}

	if memberCount == 0 {
		wow.UnregisterUpdate(updateTextures)
	}
}

func ShowUnit(ctx context.Ctx, guid wow.GUID) {
	addMember(guid)
	ctx.OnCancel(nil, func(_ context.Ctx, _ interface{}) {
		removeMember(guid)
	})
}

func doShowGroupMembers() {
	inRaid := wow.IsInRaid()
	prefix := "party"
	if inRaid {
		prefix = "raid"
	}

	num := wow.GetNumGroupMembers()
	for i := 1; i <= num; i++ {
		uid := wow.UnitID(prefix + luastrings.ToString(i))
		if !inRaid && i == num {
			uid = "player"
		}
		guid := wow.UnitGUID(uid)
		addMember(guid)
	}
}

func doHideGroupMembers() {
	for _, dm := range currMembers {
		removeMember(dm.GUID)
	}
}
