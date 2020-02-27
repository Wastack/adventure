package yaml

type YamlStateEntry struct {
	Name    string   `yaml:"name"`
	Story   string   `yaml:"story"`
	Actions []string `yaml:"actions"`
	Is_dead bool     `yaml:"is_dead"`
	Is_end  bool     `yaml:"is_end"`
}

type YamlDataSet struct {
	Version    string           `yaml:"version"`
	Start_node string           `yaml:"start_node"`
	Entries    []YamlStateEntry `yaml:"entries"`
}
