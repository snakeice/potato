package definitions

type PotatoConfig struct {
	Path       string             `yaml:"-"`
	AlwaysSudo bool               `yaml:"always_sudo"`
	Commands   map[string]Command `yaml:"commands"`
	Shell      []string           `yaml:"shell"`
	User       string             `yaml:"user"`
}
type Parameters struct {
	Default     string            `yaml:"default"`
	Description string            `yaml:"description"`
	Values      map[string]string `yaml:"values"`
}

type Command struct {
	Description string                     `yaml:"description"`
	Template    []string                   `yaml:"template"`
	Parameters  map[string]Parameters      `yaml:"parameters"`
	Fn          func(config *PotatoConfig) `yaml:"-"`
}
