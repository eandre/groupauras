package auratrack

import (
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type Target struct {
	Ctx  context.Ctx
	GUID wow.GUID
	Name string
	Self bool
	Data interface{} // Arbitrary storage for users

	cancelFunc func()
}

type Tracker struct {
	// Context that is valid for the duration of the tracker.
	// All auras tracked will be child contexts of this context.
	Ctx     context.Ctx
	Targets map[wow.GUID]*Target

	spellID int64
	cfg     *Cfg
}

type Cfg struct {
	Add    func(target *Target)
	Remove func(target *Target)
}

func New(ctx context.Ctx, spellID int64, cfg *Cfg) *Tracker {
	t := &Tracker{
		Ctx:     ctx,
		Targets: make(map[wow.GUID]*Target),
		spellID: spellID,
		cfg:     cfg,
	}

	t.enable()
	ctx.OnCancel(t, disableTracker)
	return t
}

var ongoingTrackers = make(map[int64]map[*Tracker]bool)

func disableTracker(ctx context.Ctx, data interface{}) {
	data.(*Tracker).disable()
}

func (t *Tracker) enable() {
	if ongoingTrackers[t.spellID] == nil {
		ongoingTrackers[t.spellID] = make(map[*Tracker]bool)
	}
	ongoingTrackers[t.spellID][t] = true
}

func (t *Tracker) disable() {
	delete(ongoingTrackers[t.spellID], t)
}

func init() {
	selfGUID := wow.UnitGUID("player")
	wow.RegisterEvent("COMBAT_LOG_EVENT_UNFILTERED", func(_ string, args []interface{}) {
		event := args[1].(string)
		if event != "SPELL_AURA_APPLIED" && event != "SPELL_AURA_REMOVED" {
			return
		}

		spellID := args[11].(int64)
		trackers := ongoingTrackers[spellID]
		if len(trackers) == 0 {
			return
		}
		destGUID := args[7].(wow.GUID)

		if event == "SPELL_AURA_APPLIED" {
			for t := range trackers {
				ctx, cancel := context.New(t.Ctx)
				target := &Target{
					Ctx:        ctx,
					cancelFunc: cancel,
					GUID:       destGUID,
					Name:       args[8].(string),
					Self:       destGUID == selfGUID,
				}
				t.Targets[destGUID] = target
				if t.cfg.Add != nil {
					t.cfg.Add(target)
				}
			}
			return
		}

		// SPELL_AURA_REMOVED
		for t := range trackers {
			target := t.Targets[destGUID]
			if target != nil {
				delete(t.Targets, destGUID)
				target.cancelFunc()
				if t.cfg.Remove != nil {
					t.cfg.Remove(target)
				}
			}
		}
	})
}
