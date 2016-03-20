package debug

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/core/trigger"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/shim/bridge"
)

var Point *draw.Point

func onWorldEnter(event string, args []interface{}) {
	a := aura.New("Blackhand Mines")
	//a.Enables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) print(\"enabled\"); self._count = 0; return true end"
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

	Point = draw.NewPoint(&draw.PointCfg{
		Pos:           draw.StaticPlayerPosition(),
		Texture:       "diamond",
		RotateDegrees: 360,
		RotateSpeed:   5,
	})
}

func init() {
	bridge.RegisterEvent("PLAYER_ENTERING_WORLD", onWorldEnter)
}
