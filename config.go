package groupauras

import "github.com/eandre/groupauras/core/aura"

type ProfileConfig struct {
	RotateMap bool
	Auras     []*aura.Aura
}

type Config struct {
	profile ProfileConfig
}

var defaultConfig = &Config{
	profile: ProfileConfig{
		RotateMap: true,
		Auras:     []*aura.Aura{},
	},
}
