package debug

import (
	"github.com/eandre/groupauras/aura"
	"github.com/eandre/groupauras/bridge"
	"github.com/eandre/groupauras/trigger"
)

func OnWorldEnter(event string, args []interface{}) {
	a := aura.New("Blackhand Mines")
	a.Enables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) print(\"enabled\"); self._count = 0; return true end"
	a.Disables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) return self._count > 10 end"
	a.Events["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) self._count = self._count + 1; print(self.Aura.Name, event, unpack(args)) end"
	ca, err := aura.Compile(a)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Compiled aura:", ca.Aura.Name)

	t := trigger.NewWatcher()
	t.Add(ca)
}

func init() {
	bridge.RegisterEvent("PLAYER_ENTERING_WORLD", OnWorldEnter)
}
