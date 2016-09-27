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

var groupMemberSources = 0

func ShowGroupMembers(ctx context.Ctx) {
	groupMemberSources += 1
	doShowGroupMembers()

	ctx.OnCancel(nil, func(c context.Ctx, _ interface{}) {
		groupMemberSources -= 1
		if groupMemberSources <= 0 {
			doHideGroupMembers()
			groupMemberSources = 0
		}
	})
}

type drawingMember struct {
	Point *draw.Point
	GUID  wow.GUID
}

var currMembers []*drawingMember

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

func doShowGroupMembers() {
	wow.RegisterUpdate(updateTextures)
	if wow.IsInRaid() {
		num := wow.GetNumGroupMembers()
		for i := 1; i <= num; i++ {
			uid := wow.UnitID("raid" + luastrings.ToString(i))
			r, g, b := classutil.Color(uid)
			pt := draw.NewPoint(&draw.PointCfg{
				Pos:         draw.UnitPosition(uid),
				Texture:     "party",
				VertexColor: []float32{r, g, b, 1},
				SizeYards:   1,
				DrawLayer:   widget.LayerOverlay,
			})
			currMembers = append(currMembers, &drawingMember{
				Point: pt,
				GUID:  wow.UnitGUID(uid),
			})
		}
	} else {
		// If we're not in a raid group, still draw the player
		r, g, b := classutil.Color("player")
		pt := draw.NewPoint(&draw.PointCfg{
			Pos:         draw.UnitPosition("player"),
			Texture:     "party",
			VertexColor: []float32{r, g, b, 1},
			SizeYards:   1,
			DrawLayer:   widget.LayerOverlay,
		})
		currMembers = append(currMembers, &drawingMember{
			Point: pt,
			GUID:  wow.UnitGUID("player"),
		})
	}
}

func doHideGroupMembers() {
	wow.UnregisterUpdate(updateTextures)
	for _, dm := range currMembers {
		dm.Point.Free(false)
	}
	currMembers = []*drawingMember{}
}
