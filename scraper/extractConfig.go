package scraper

import "github.com/PuerkitoBio/goquery"

func ExtractConfig(sel *goquery.Selection, fields map[string]FieldConfig) map[string]interface{} {
	result := make(map[string]interface{})
	for key, field := range fields {
		result[key] = field.Extract(sel)
	}
	return result
}
