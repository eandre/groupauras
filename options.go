package groupauras

import "github.com/eandre/groupauras/pkg/ace/acecfg"

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

func init() {
	acecfg.RegisterOptionsTable("GroupAuras", CoreOptions, []string{"groupauras", "ga"})
	acecfg.AddToBlizOptions("GroupAuras", "")
}
