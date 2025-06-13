package scraper

type Transform []string

type Config struct {
	Selector string                 `yaml:"selector"`
	Fields   map[string]FieldConfig `yaml:"fields"`
}

type FieldConfig struct {
	Type      string                 `yaml:"type"`
	Selector  string                 `yaml:"selector"`
	Transform interface{}            `yaml:"transform"`
	Item      *FieldConfig           `yaml:"item"`
	Fields    map[string]FieldConfig `yaml:"fields"`
	Constant  interface{}            `yaml:"constant"`
	Union     []FieldConfig          `yaml:"union"`
}

type ConfigWithSelector struct {
	Selector string
}

type UnionConfig struct {
	Configs []FieldConfig
}
