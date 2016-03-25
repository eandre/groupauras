package debug

import (
	"github.com/eandre/groupauras/core/aura"
	"github.com/eandre/groupauras/core/trigger"
	"github.com/eandre/groupauras/pkg/ace/acecfg"
	"github.com/eandre/groupauras/pkg/ace/acedb"
	"github.com/eandre/groupauras/pkg/draw"
	"github.com/eandre/groupauras/shim/bridge"
)

type profileConfig struct {
	RotateMap bool
}

type Config struct {
	profile profileConfig
}

var defaultConfig = &Config{
	profile: profileConfig{
		RotateMap: true,
	},
}

var DB *Config

var CoreOptions = &acecfg.Group{
	Name: "General",
	Children: map[string]acecfg.Option{
		"general": &acecfg.Group{
			Name: "General Options",
			Children: map[string]acecfg.Option{
				"rotateMap": &acecfg.Toggle{
					Name: "Rotate Map",
					Desc: "Rotates the map around you as you move.",
					Get:  func(*acecfg.Info) bool { return DB.profile.RotateMap },
					Set:  func(info *acecfg.Info, value bool) { DB.profile.RotateMap = value },
				},
			},
		},
	},
}

var Point *draw.Point

func onWorldEnter(event string, args []interface{}) {
	DB = acedb.New("GroupAurasDB", defaultConfig, "").(*Config)
	acecfg.RegisterOptionsTable("GroupAuras", CoreOptions, []string{"groupauras", "ga"})
	acecfg.AddToBlizOptions("GroupAuras", "")

	a := aura.New("Blackhand Mines")
	//a.Enables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) print(\"enabled\"); self._count = 0; return true end"
	a.Disables["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) return self._count > 10 end"
	a.Events["COMBAT_LOG_EVENT_UNFILTERED"] = "function(self, event, args) self._count = self._count + 1; print(self.Aura.Name, event, unpack(args)) end"
	ca, err := aura.Compile(a)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	t := trigger.NewWatcher()
	t.Add(ca)

	Point = draw.NewPoint(&draw.PointCfg{
		Pos:           draw.StaticPlayerPosition(),
		Texture:       "timer",
		SizeYards:     40,
		RotateDegrees: 360,
		RotateSpeed:   5,
	})

	draw.NewPoint(&draw.PointCfg{
		Pos:        draw.PlayerPosition(),
		Texture:    "diamond",
		SizePixels: 5,
	})
}

func init() {
	bridge.RegisterEvent("PLAYER_ENTERING_WORLD", onWorldEnter)
}
