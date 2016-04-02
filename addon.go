package groupauras

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/pkg/ace/aceaddon"
	"github.com/eandre/groupauras/pkg/ace/acedb"
)

var DB *Config

type Addon struct{}

func (a *Addon) OnInitialize() {
	DB = acedb.New("GroupAurasDB", defaultConfig, "").(*Config)
	debug()
}

func init() {
	aceaddon.New("GroupAuras", &Addon{})
}

func debug() {
	a := aura.New("Blackhand Mines")
	//a.Enables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) print(\"enabled\"); self._count = 0; return true end"
	a.Disables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) return self._count > 10 end"
	a.Events["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) self._count = self._count + 1; print(self.Aura.Name, event, unpack(args)) end"
	_, err := aura.Compile(a)
	if err != nil {
		println("Error:", err.Error())
		return
	}
}
