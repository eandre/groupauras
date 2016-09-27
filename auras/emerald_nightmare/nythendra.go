package emerald_nightmare

import (
	"github.com/eandre/groupauras/pkg/combatevent"
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/groupauras/pkg/raidutil/rdu"
	"github.com/eandre/lunar-shim/hbd"
	"github.com/eandre/lunar-wow/pkg/wow"
)

const (
	RotRange    = 8
	RotDuration = 9
)

type Rot struct {
	Target wow.GUID
	Name   string
	Self   bool
	Point  *draw.Point
}

func (r *Rot) Free() {
	if r.Self {
		rdu.HideGroupMembers("rot")
	}
	r.Point.Free(false)
}

func (r *Rot) updateColor() {
	alert := false
	if r.Self {
		// If we're the target, nobody else should be in range
		alert = len(raidutil.GroupMembersWithin(r.Target, RotRange)) > 0
	} else {
		// If we're not the target, we only care about ourselves being out of range
		uid, ok := raidutil.GUIDToUnitID(r.Target)
		if ok {
			px, py, pi := hbd.PlayerWorldPosition()
			tx, ty, ti := hbd.UnitWorldPosition(uid)
			if pi == ti {
				dist := hbd.WorldDistance(pi, px, py, tx, ty)
				alert = dist <= RotRange
			}
		}
	}

	if alert {
		r.Point.SetVertexColor(1, 0, 0, 1)
		r.Point.SetText("|cFFFF0000" + r.Name + "|r")
	} else {
		r.Point.SetVertexColor(1, 1, 1, 0.5)
		r.Point.SetText(r.Name)
	}
}

var currTargets = make(map[wow.GUID]*Rot)

var elapsed float32

func updateRangeDisplay(dt float32) {
	elapsed += dt
	if elapsed < 0.1 {
		return
	}
	elapsed -= 0.1
	for _, r := range currTargets {
		r.updateColor()
	}
}

func enableDisplay() {
	wow.RegisterUpdate(updateRangeDisplay)
}

func disableDisplay() {
	wow.UnregisterUpdate(updateRangeDisplay)
	elapsed = 0
}

func spellAuraApplied(sourceName, destName string, sourceGUID, destGUID wow.GUID, spellID int64, auraType string, amount float32) {
	if spellID != 139 {
		return
	}

	if currTargets[destGUID] != nil {
		currTargets[destGUID].Free()
	}

	uid, ok := raidutil.GUIDToUnitID(destGUID)
	if !ok {
		return
	}

	name, _ := wow.UnitName(uid)
	r := &Rot{
		Target: destGUID,
		Name:   name,
		Self:   wow.UnitIsUnit(uid, "player"),
		Point: draw.NewPoint(&draw.PointCfg{
			Pos:           draw.UnitPosition(uid),
			Text:          name,
			Texture:       "timer",
			SizeYards:     RotRange,
			RotateDegrees: 360,
			RotateSpeed:   RotDuration,
		}),
	}

	if r.Self {
		rdu.ShowGroupMembers("rot")
	}

	currTargets[destGUID] = r
	r.updateColor()
	enableDisplay()
}

func spellAuraRemoved(sourceName, destName string, sourceGUID, destGUID wow.GUID, spellID int64, auraType string, amount float32) {
	if spellID != 139 {
		return
	}

	r := currTargets[destGUID]
	if r != nil {
		r.Free()
		delete(currTargets, destGUID)
	}

	if len(currTargets) == 0 {
		disableDisplay()
	}
}

func init() {
	combatevent.OnSpellAuraApplied(context.Base, spellAuraApplied)
	combatevent.OnSpellAuraRemoved(context.Base, spellAuraRemoved)
}
