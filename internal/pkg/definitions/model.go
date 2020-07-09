package definitions

type PotatoConfig struct {
	Path       string
	AlwaysSudo bool               `yaml:"always_sudo"`
	Commands   map[string]Command `yaml:"commands"`
	Shell      string             `yaml:"shell"`
	User       string             `yaml:"user"`
	Editor     string             `yaml:"editor"`
}
type Parameters struct {
	Default     string   `yaml:"default"`
	Description string   `yaml:"description"`
	Values      []string `yaml:"values"`
}

type Command struct {
	Description string                `yaml:"description"`
	Template    string                `yaml:"template"`
	Parameters  map[string]Parameters `yaml:"parameters"`
}
