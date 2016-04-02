package groupauras

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
