package scraper

import (
	"gopkg.in/yaml.v3"
)

func ParseDynamicYAML(yamlStr string) (Config, error) {
	var raw map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &raw)
	if err != nil {
		return Config{}, err
	}

	fields := make(map[string]FieldConfig)
	if f, ok := raw["fields"]; ok {
		fieldMap := f.(map[string]interface{})
		for k, v := range fieldMap {
			fields[k] = parseFieldConfig(v.(map[string]interface{}))
		}
	}

	return Config{
		Selector: raw["selector"].(string),
		Fields:   fields,
	}, nil
}

func parseFieldConfig(data map[string]interface{}) FieldConfig {
	fc := FieldConfig{}

	if t, ok := data["type"]; ok {
		fc.Type = t.(string)
	}
	if s, ok := data["selector"]; ok {
		fc.Selector = s.(string)
	}
	if tr, ok := data["transform"]; ok {
		fc.Transform = tr
	}
	if c, ok := data["constant"]; ok {
		fc.Constant = c
	}
	if item, ok := data["item"]; ok {
		itemMap := item.(map[string]interface{})
		itemConfig := parseFieldConfig(itemMap)
		fc.Item = &itemConfig
	}
	if fields, ok := data["fields"]; ok {
		fieldMap := fields.(map[string]interface{})
		fc.Fields = make(map[string]FieldConfig)
		for k, v := range fieldMap {
			fc.Fields[k] = parseFieldConfig(v.(map[string]interface{}))
		}
	}
	if union, ok := data["union"]; ok {
		unionList := union.([]interface{})
		fc.Union = make([]FieldConfig, 0)
		for _, u := range unionList {
			fc.Union = append(fc.Union, parseFieldConfig(u.(map[string]interface{})))
		}
	}

	return fc
}
