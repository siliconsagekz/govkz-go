package config

type Config struct {
	IsDebug 	*bool `yaml:"is_debug"`
	Services 	StateServices `yaml:"services"`
}

type StateServices struct {
	External	StateExternal `yaml:"external"`
}
type StateExternal struct {
	Stat_gov 	string `yaml:"stat_gov"`
}