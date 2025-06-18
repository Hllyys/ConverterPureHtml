package scraper

import (
	"gopkg.in/yaml.v3"
)

func ParseYAMLFromString(yamlStr string) (Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	return config, err
}
