package groupauras

import "github.com/eandre/groupauras/pkg/ace/acecfg"

var CoreOptions *acecfg.Group

func initOptions() {
	CoreOptions = &acecfg.Group{
		Name: "General",
		Children: map[string]acecfg.Option{
			"general": &acecfg.Group{
				Name:  "General Options",
				Order: 1,

				Children: map[string]acecfg.Option{
					"rotateMap": &acecfg.Toggle{
						Name: "Rotate Map",
						Desc: "Rotates the map around you as you move.",
						Get:  func(*acecfg.Info) bool { return DB.profile.RotateMap },
						Set:  func(info *acecfg.Info, value bool) { DB.profile.RotateMap = value },
					},
				},
			},

			"auras": &acecfg.Group{
				Name:  "Auras",
				Order: 2,

				Children: map[string]acecfg.Option{
					"addAura": &acecfg.Button{
						Name: "Add aura",
						Func: func(info *acecfg.Info) { addAura(info) },
					},
					"auras": &acecfg.Group{
						Name: "Auras",

						Children: map[string]acecfg.Option{},
					},
				},
			},
		},
	}
}

func addAura(info *acecfg.Info) {
	//aura := Library.Add("Unnamed Aura")
	//children := CoreOptions.Children["auras"].(*acecfg.Group).Children["auras"].(*acecfg.Group).Children
	//children[aura.ID] = &acecfg.Group{
	//	Name: aura.Name,
	//	Children: map[string]acecfg.Option{
	//		"enable": &acecfg.Toggle{
	//			Name:  "Enable",
	//			Order: 1,
	//		},
	//	},
	//}
}

func init() {
	initOptions()
	acecfg.RegisterOptionsTable("GroupAuras", CoreOptions, []string{"groupauras", "ga"})
	acecfg.AddToBlizOptions("GroupAuras", "")
}
