package scraper

import "github.com/PuerkitoBio/goquery"

func (u *UnionConfig) Extract(sel *goquery.Selection) interface{} {
	for _, config := range u.Configs {
		result := config.Extract(sel)
		if result != nil {
			return result
		}
	}
	return nil
}
