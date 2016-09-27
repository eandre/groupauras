package hfc

import (
	"github.com/eandre/groupauras/pkg/combatevent"
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/groupauras/pkg/raidutil/rdu"
	"github.com/eandre/lunar-wow/pkg/wow"
)

const (
	sfRange      = 8
	sfMinSoakers = 1
	sfDuration   = 5
)

type ShadowfelBurst struct {
	Target wow.GUID
	Point  *draw.Point
}

func (sf *ShadowfelBurst) Free() {
	sf.Point.Free(false)
}

func (sf *ShadowfelBurst) EnoughSoakers() bool {
	return len(raidutil.GroupMembersWithin(sf.Target, sfRange)) >= sfMinSoakers
}

func (sf *ShadowfelBurst) updateColor() {
	if sf.EnoughSoakers() {
		sf.Point.SetVertexColor(0, 1, 0, 1)
	} else {
		sf.Point.SetVertexColor(1, 0, 0, 1)
	}
}

var currTargets = make(map[wow.GUID]*ShadowfelBurst)

var elapsed float32

func updateRangeDisplay(dt float32) {
	elapsed += dt
	if elapsed < 0.1 {
		return
	}
	elapsed -= 0.1
	for _, sf := range currTargets {
		sf.updateColor()
	}
}

func enableDisplay() {
	rdu.ShowGroupMembers("shadowfel-burst")
	wow.RegisterUpdate(updateRangeDisplay)
}

func disableDisplay() {
	rdu.HideGroupMembers("shadowfel-burst")
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

	sf := &ShadowfelBurst{
		Target: destGUID,
		Point: draw.NewPoint(&draw.PointCfg{
			Pos:           draw.StaticUnitPosition(uid),
			Texture:       "timer",
			SizeYards:     sfRange,
			RotateDegrees: 360,
			RotateSpeed:   sfDuration,
		}),
	}
	currTargets[destGUID] = sf
	sf.updateColor()
	enableDisplay()
}

func spellAuraRemoved(sourceName, destName string, sourceGUID, destGUID wow.GUID, spellID int64, auraType string, amount float32) {
	if spellID != 139 {
		return
	}

	sf := currTargets[destGUID]
	if sf != nil {
		sf.Free()
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
