package scraper

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractWithConfig(html string, cfg Config, url string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	CheckErr(err)

	var results []interface{}

	doc.Find(cfg.Selector).Each(func(i int, s *goquery.Selection) {
		val := ExtractConfig(s, cfg.Fields)
		val["source_url"] = url
		results = append(results, val)
	})

	data, err := json.MarshalIndent(results, "", "  ")
	CheckErr(err)

	return string(data), nil
}
