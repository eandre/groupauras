package groupauras

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/core/library"
	"github.com/eandre/groupauras/core/library/transfer"
	"github.com/eandre/groupauras/core/runner"
	"github.com/eandre/groupauras/core/watcher"
	"github.com/eandre/groupauras/pkg/ace/acedb"
	"github.com/eandre/lunar-shim/ace/aceaddon"
)

// import _ "github.com/eandre/groupauras/auras/emerald_nightmare"
import _ "github.com/eandre/groupauras/debug"
import _ "github.com/eandre/groupauras/pkg/raidutil/inspect"

var (
	DB      *Config
	Library *library.Library
	Runner  *runner.Runner
	Watcher *watcher.Watcher
)

type Addon struct{}

func (a *Addon) OnInitialize() {
	DB = acedb.New("GroupAurasDB", defaultConfig, "").(*Config)
	Library = library.New(DB.profile.Auras)
	Runner = runner.New()
	Watcher = watcher.New(Library, Runner)
}

func init() {
	aceaddon.New("GroupAuras", &Addon{})
	a := &aura.Aura{
		ID:       "test-id",
		Revision: 1,
		Sig:      "sig",
		Raw:      "raw",
	}
	transfer.Send(a, "WHISPER", "Swedeheart", nil, nil)
}

func debug() {
	//a := aura.New("Blackhand Mines")
	//a.Enables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) print(\"enabled\"); self._count = 0; return true end"
	//a.Disables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) return self._count > 10 end"
	//a.Events["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) self._count = self._count + 1; print(self.Aura.Name, event, unpack(args)) end"
	//_, err := aura.Compile(a)
	//if err != nil {
	//	println("Error:", err.Error())
	//	return
	//}
}
