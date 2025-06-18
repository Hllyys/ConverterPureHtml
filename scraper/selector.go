package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

func (c *ConfigWithSelector) GetAllMatches(sel *goquery.Selection, includeRoot bool) *goquery.Selection {
	if c.Selector == "" {
		return sel
	}

	if includeRoot && sel.Is(c.Selector) {
		return sel
	}

	return sel.Find(c.Selector)
}

func (c *ConfigWithSelector) GetFirstMatch(sel *goquery.Selection, includeRoot bool) *goquery.Selection {
	return c.GetAllMatches(sel, includeRoot).First()
}
