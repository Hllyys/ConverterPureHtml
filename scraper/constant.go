package scraper

import "github.com/PuerkitoBio/goquery"

func (c *FieldConfig) Extract(sel *goquery.Selection) interface{} {
	if c.Selector != "" {
		sel = sel.Find(c.Selector)
		if sel.Length() == 0 {
			return nil
		}
	}
	return c.Constant
}
