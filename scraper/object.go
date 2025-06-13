package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractObject(field FieldConfig, sel *goquery.Selection) interface{} {

	config := ConfigWithSelector{Selector: field.Selector}
	sel = config.GetFirstMatch(sel, false)

	if field.Fields == nil {
		return map[string]interface{}{}
	}

	result := make(map[string]interface{})

	for key, subField := range field.Fields {
		switch subField.Type {
		case "primitive":
			result[key] = ExtractPrimitive(subField, sel)
		case "array":
			result[key] = ExtractArray(subField, sel)
		case "object":
			result[key] = ExtractObject(subField, sel)
		case "constant":
			result[key] = subField.Constant
		case "union":
			unionConfig := UnionConfig{Configs: subField.Union}
			result[key] = unionConfig.Extract(sel)
		default:
			result[key] = nil
		}
	}

	return result
}
