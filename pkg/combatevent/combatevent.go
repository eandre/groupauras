package combatevent

import (
	"github.com/eandre/groupauras/pkg/context"
	"github.com/eandre/lunar-wow/pkg/wow"
)

type SpellAuraApplied func(sourceName, destName string, sourceGUID, destGUID wow.GUID, spellID int64, auraType string, amount float32)
type SpellAuraRemoved func(sourceName, destName string, sourceGUID, destGUID wow.GUID, spellID int64, auraType string, amount float32)

var (
	spellAuraApplied = make(map[interface{}]int)
	spellAuraRemoved = make(map[interface{}]int)
)

func OnSpellAuraApplied(ctx context.Ctx, f SpellAuraApplied) { handleMap(ctx, spellAuraApplied, f) }
func OnSpellAuraRemoved(ctx context.Ctx, f SpellAuraRemoved) { handleMap(ctx, spellAuraRemoved, f) }

func handleMap(ctx context.Ctx, m map[interface{}]int, key interface{}) {
	m[key] = m[key] + 1
	ctx.OnCancel(nil, func(context.Ctx, interface{}) {
		n := m[key] - 1
		if n <= 0 {
			delete(m, key)
		} else {
			m[key] = n
		}
	})
}

func onCombatLogEvent(event string, args []interface{}) {
	var (
		subEvent   = args[1].(string)
		sourceGUID = args[3].(wow.GUID)
		sourceName = args[4].(string)
		destGUID   = args[7].(wow.GUID)
		destName   = args[8].(string)
	)

	if subEvent == "SPELL_AURA_APPLIED" {
		spellID := args[11].(int64)
		auraType := args[14].(string)
		amount := args[15].(float32)
		for f := range spellAuraApplied {
			f.(SpellAuraApplied)(sourceName, destName, sourceGUID, destGUID, spellID, auraType, amount)
		}
	} else if subEvent == "SPELL_AURA_REMOVED" {
		spellID := args[11].(int64)
		auraType := args[14].(string)
		amount := args[15].(float32)
		for f := range spellAuraRemoved {
			f.(SpellAuraRemoved)(sourceName, destName, sourceGUID, destGUID, spellID, auraType, amount)
		}
	}
}

func init() {
	wow.RegisterEvent("COMBAT_LOG_EVENT_UNFILTERED", onCombatLogEvent)
}
