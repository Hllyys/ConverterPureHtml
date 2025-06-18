package scraper

import "github.com/PuerkitoBio/goquery"

func ExtractObject(field FieldConfig, sel *goquery.Selection) interface{} {
	base := BaseExtractor{Selector: field.Selector}
	selected := base.ApplySelector(sel).First()

	if selected == nil {
		return nil
	}

	return ExtractConfig(selected, field.Fields)
}
