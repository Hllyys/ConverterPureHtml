package scraper

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func ExtractPrimitive(field FieldConfig, sel *goquery.Selection) interface{} {
	base := BaseExtractor{Selector: field.Selector}
	selected := base.ApplySelector(sel).First()

	if selected == nil {
		return nil
	}

	val := selected.Text()

	transforms := NormalizeTransform(field.Transform)
	if transforms == nil {
		return strings.TrimSpace(val)
	}

	return ApplyTransform(val, selected, transforms)
}
