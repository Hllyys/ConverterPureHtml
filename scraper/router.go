package scraper

import "github.com/PuerkitoBio/goquery"

func (field *FieldConfig) Extract(sel *goquery.Selection) interface{} {
	switch field.Type {
	case "primitive":
		return ExtractPrimitive(*field, sel)
	case "object":
		return ExtractObject(*field, sel)
	case "array":
		return ExtractArray(*field, sel)
	case "constant":
		return field.Constant
	case "union":
		unionConfig := UnionConfig{Configs: field.Union}
		return unionConfig.Extract(sel)
	default:
		return nil
	}
}
