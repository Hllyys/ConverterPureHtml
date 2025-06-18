package scraper

import "github.com/PuerkitoBio/goquery"

func ExtractArray(field FieldConfig, sel *goquery.Selection) interface{} {
	var results []interface{}
	base := BaseExtractor{Selector: field.Selector}
	matches := base.ApplySelector(sel)

	if matches.Length() == 0 {
		return results
	}

	var itemConfig *FieldConfig
	if field.Item != nil {
		itemConfig = field.Item
	} else {
		itemConfig = &FieldConfig{
			Type:      "primitive",
			Selector:  "",
			Transform: field.Transform,
		}
	}

	matches.Each(func(i int, s *goquery.Selection) {
		results = append(results, itemConfig.Extract(s))
	})

	return results
}
