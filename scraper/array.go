package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractArray(field FieldConfig, sel *goquery.Selection) interface{} {
	var results []interface{}
	base := BaseExtractor{Selector: field.Selector}
	matches := base.ApplySelector(sel)

	if matches.Length() == 0 {
		return results
	}

	// Tüm eşleşen elemanları sırayla işle
	matches.Each(func(i int, s *goquery.Selection) {
		itemConfig := field.Item
		if itemConfig == nil {
			itemConfig = &FieldConfig{
				Type:      "primitive",
				Selector:  "",
				Transform: field.Transform,
			}
		}

		var val interface{}
		switch itemConfig.Type {
		case "primitive":
			val = ExtractPrimitive(*itemConfig, s)
		case "object":
			val = ExtractObject(*itemConfig, s)
		default:
			val = nil
		}

		results = append(results, val)
	})

	return results
}
