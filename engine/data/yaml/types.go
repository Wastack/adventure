package yaml

type YamlGameAction struct {
	Target string `yaml:"target"`
	Name   string `yaml:"name"`
	Story  string `yaml:"story"`
	Secret string `yaml:"secret"`
}

type YamlStateEntry struct {
	Name    string           `yaml:"name"`
	Story   string           `yaml:"story"`
	Actions []YamlGameAction `yaml:"actions"`
	Is_dead bool             `yaml:"is_dead"`
	Is_end  bool             `yaml:"is_end"`
}

type YamlDataSet struct {
	Version    string           `yaml:"version"`
	Start_node string           `yaml:"start_node"`
	Entries    []YamlStateEntry `yaml:"entries"`
}
