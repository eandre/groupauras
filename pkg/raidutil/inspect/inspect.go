package inspect

import (
	"github.com/eandre/groupauras/pkg/raidutil"
	"github.com/eandre/lunar-wow/pkg/luastrings"
	"github.com/eandre/lunar-wow/pkg/widget"
	"github.com/eandre/lunar-wow/pkg/wow"
	"github.com/eandre/lunar/lua"
)

type Data struct {
	InspectTime wow.Time
	SpecID      wow.SpecID
}

var lastInspectPerGUID = make(map[wow.GUID]*Data)
var lastInspect wow.Time

func DataForGUID(guid wow.GUID) (data *Data, ok bool) {
	data = lastInspectPerGUID[guid]
	return data, data != nil
}

func onInspectReady(event string, args []interface{}) {
	guid := args[0].(wow.GUID)
	now := wow.GetTime()
	lastInspect = now

	// Prevent multiple inspections in a short timespan due to event spam
	if last := lastInspectPerGUID[guid]; last != nil && (now-last.InspectTime) < 0.2 {
		return
	}

	data := &Data{
		InspectTime: now,
	}
	lastInspectPerGUID[guid] = data

	// Get the unit id of the inspected player
	unitID, ok := raidutil.GUIDToUnitID(guid)
	if !ok {
		return
	}

	// Get the specialization of the player
	specID := wow.GetInspectSpecialization(unitID)
	if specID == 0 {
		// Couldn't find spec; skip
		return
	}

	data.SpecID = specID
}

func init() {
	wow.RegisterEvent("INSPECT_READY", onInspectReady)
}

var inspectQueue = make(map[wow.GUID]wow.Time) // guid -> time added to queue

func nextInspect() {
	// If we're in combat or the inspect frame is visible, don't inspect
	if wow.IsEncounterInProgress() {
		return
	}
	inspectFrame := lua.Raw(`InspectFrame`).(widget.Frame)
	if inspectFrame != nil && inspectFrame.IsShown() {
		return
	}

	var unqueue []wow.GUID
	now := wow.GetTime()
	for guid, added := range inspectQueue {
		unitID, ok := raidutil.GUIDToUnitID(guid)
		if !ok {
			unqueue = append(unqueue, guid)
		} else if wow.CanInspect(unitID) {
			wow.NotifyInspect(unitID)
			unqueue = append(unqueue, guid)
		} else if added < (now - 300) {
			unqueue = append(unqueue, guid)
		}
	}

	for _, guid := range unqueue {
		delete(inspectQueue, guid)
	}
}

func queueGroupInspect() {
	prefix := "party"
	inRaid := wow.IsInRaid()
	if inRaid {
		prefix = "raid"
	}

	now := wow.GetTime()
	num := wow.GetNumGroupMembers()
	for i := 1; i <= num; i++ {
		online := true
		unitID := wow.UnitID(prefix + luastrings.ToString(i))
		if inRaid {
			// If we're in a raid, check if players are online.
			// If not, we only have max 5 players to check, so
			// we can afford to skip it.
			_, _, _, _, _, _, _, online, _, _, _ = wow.GetRaidRosterInfo(i)
		} else {
			// In party; if it's the last index, unit ID is "player", not "partyN"
			if i == num {
				unitID = "player"
			}
		}
		if online {
			guid := wow.UnitGUID(unitID)
			inspectQueue[guid] = now
		}
	}
}

var timer float32 = -5
var queueTimer float32 = 0

func timerUpdate(dt float32) {
	timer += dt
	if timer >= 3.5 {
		queueTimer += timer
		timer = 0
		if queueTimer > 60 {
			queueTimer = 0
			queueGroupInspect()
		}
		nextInspect()
	}
}

func init() {
	wow.RegisterUpdate(timerUpdate)
}
