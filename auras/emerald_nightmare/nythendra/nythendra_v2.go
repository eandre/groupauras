package nythendra

import (
	"github.com/eandre/groupauras/pkg/classutil"
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/groupauras/pkg/raidutil/auratrack"
	"github.com/eandre/groupauras/pkg/raidutil/inspect"
	"github.com/eandre/groupauras/pkg/raidutil/rdu"
	"github.com/eandre/lunar-wow/pkg/wow"
)

const (
	InfestedMindID             = 61295 // Riptide, real: 205043
	InfestedMindInterrupters   = 2     // need 2 interrupters
	InfestedMindInterruptRange = 10
)

type Nythendra struct {
	Ctx          context.Ctx
	InfestedMind *auratrack.Tracker
	dt           float32
}

func (n *Nythendra) updateInfestedMind(dt float32) {
	n.dt += dt
	if n.dt < 0.1 {
		return
	}
	n.dt -= 0.1

	for _, target := range n.InfestedMind.Targets {
		close := raidutil.GroupMembersWithin(target.GUID, InfestedMindInterruptRange)
		interrupters := 0
		for _, c := range close {
			// Make sure the person isn't MC'd
			if n.InfestedMind.Targets[c.GUID] == nil {
				if hasInterrupt(c.GUID) {
					interrupters += 1
				}
			}
		}
		pt := target.Data.(*draw.Point)
		if interrupters >= InfestedMindInterrupters {
			pt.SetVertexColor(0.25, 1, 0.25, 1)
		} else {
			pt.SetVertexColor(1, 0.25, 0.25, 1)
		}
	}
}

func Enable(ctx context.Ctx) *Nythendra {
	n := &Nythendra{Ctx: ctx}
	n.InfestedMind = auratrack.New(ctx, InfestedMindID, &auratrack.Cfg{
		Add: func(target *auratrack.Target) {
			pt := draw.NewPoint(&draw.PointCfg{
				Ctx:       target.Ctx,
				Pos:       draw.GUIDPosition(target.GUID),
				Text:      target.Name,
				Texture:   "highlight",
				SizeYards: InfestedMindInterruptRange,
			})
			target.Data = pt
			rdu.ShowGroupMembers(target.Ctx)
			if len(n.InfestedMind.Targets) == 1 {
				wow.RegisterUpdate(n.updateInfestedMind)
			}
		},
		Remove: func(target *auratrack.Target) {
			if len(n.InfestedMind.Targets) == 0 {
				wow.UnregisterUpdate(n.updateInfestedMind)
			}
		},
	})

	return n
}

func hasInterrupt(guid wow.GUID) bool {
	data, ok := inspect.DataForGUID(guid)
	if !ok {
		return false
	}
	spec := classutil.Specs[data.SpecID]
	if spec == nil {
		println("groupauras: No spec found for id", data.SpecID)
		return false
	}
	return spec.Melee && spec.HasInterrupt
}
