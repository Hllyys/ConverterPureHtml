package scraper

import "github.com/PuerkitoBio/goquery"

func ExtractConfig(sel *goquery.Selection, fields map[string]FieldConfig) map[string]interface{} {
	result := make(map[string]interface{})

	for key, field := range fields {
		switch field.Type {

		case "primitive":
			result[key] = ExtractPrimitive(field, sel)
		case "array":
			result[key] = ExtractArray(field, sel)
		case "object":
			result[key] = ExtractObject(field, sel)
		case "constant":
			result[key] = field.Constant
		case "union":
			unionConfig := UnionConfig{Configs: field.Union}
			result[key] = unionConfig.Extract(sel)
		default:
			result[key] = nil
		}
	}

	return result
}
