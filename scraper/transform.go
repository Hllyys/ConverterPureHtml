package scraper

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Her transform adımında dönüşüm sonucu value'ya atanır.
func ApplyTransform(value string, sel *goquery.Selection, transforms Transform) string {
	for _, t := range transforms {
		value = applySingleTransform(value, sel, t)
	}
	return value
}

// Tek bir transform işlemini uygular.
func applySingleTransform(val string, sel *goquery.Selection, t string) string {
	switch {
	case t == "trim":
		return strings.TrimSpace(val)

	case t == "date":
		parts := strings.Split(val, "-")
		if len(parts) == 3 {
			return fmt.Sprintf("%s.%s.%s", parts[2], parts[1], parts[0])
		}
		return val

	// attr(x): ilgili HTML elementinden attribute değeri okur.
	// Örn: "attr(src)" → <img src="...">
	case strings.HasPrefix(t, "attr(") && strings.HasSuffix(t, ")"):
		attr := t[5 : len(t)-1] // Parantez içindeki attribute adını alır.
		if attrVal, exists := sel.Attr(attr); exists {
			return attrVal
		}
	}

	return val
}
