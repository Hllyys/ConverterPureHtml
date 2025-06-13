package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Bu fonksiyon temel değerleri (string/text, attribute, simple value) çıkarır ve varsa transform uygular.
func ExtractPrimitive(field FieldConfig, sel *goquery.Selection) interface{} {

	config := ConfigWithSelector{Selector: field.Selector}
	selected := config.GetFirstMatch(sel, false)

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
