package types

type Cyaml map[string]interface{}

type CyamlEntry struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

type CyamlMeta struct {
	Type    string       `yaml:"name"`
	Version string       `yaml:"version"`
	Entries []CyamlEntry `yaml:"entries"`
}
