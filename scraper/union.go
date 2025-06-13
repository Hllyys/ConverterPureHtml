package scraper

import (
	"github.com/PuerkitoBio/goquery"
)


//sırayla tüm Union içindeki FieldConfig'leri dener.
func (u *UnionConfig) Extract(sel *goquery.Selection) interface{} {
	for _, config := range u.Configs {
		var result interface{}

		switch config.Type {
		case "primitive":
			result = ExtractPrimitive(config, sel)
		case "object":
			result = ExtractObject(config, sel)
		case "array":
			result = ExtractArray(config, sel)
		case "constant":
			result = config.Constant
		case "union":
			// Eğer iç içe union varsa, recursive yapar.
			nested := UnionConfig{Configs: config.Union}
			result = nested.Extract(sel)
		default:
			result = nil
		}

		if result != nil {
			return result
		}
	}
	return nil
}
