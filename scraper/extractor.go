package scraper

import "github.com/PuerkitoBio/goquery"

// Bu bizim basit "abstract class" görevindeki yapı
type BaseExtractor struct {
	Selector string
}

func (b *BaseExtractor) ApplySelector(sel *goquery.Selection) *goquery.Selection {
	if b.Selector == "" {
		return sel
	}
	return sel.Find(b.Selector)
}
