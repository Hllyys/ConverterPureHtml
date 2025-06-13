package scraper

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

func Extract(html, yamlStr, url string) (string, error) {
	var cfg Config

	// YAML'den config'i oku ve parse et
	decoder := yaml.NewDecoder(strings.NewReader(yamlStr))
	err := decoder.Decode(&cfg)
	CheckErr(err)

	// HTML içeriğini parse ederek GoQuery DOM belgesi oluştur
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	CheckErr(err)
	var results []interface{}

	// Örn: ".product" gibi tüm ürünleri bul
	doc.Find(cfg.Selector).Each(func(i int, s *goquery.Selection) {
		// Her bulunan değişkenin ExtractConfig ile detaylı alanları çıkar
		val := ExtractConfig(s, cfg.Fields)

		val["source_url"] = url

		results = append(results, val)
	})

	data, err := json.MarshalIndent(results, "", "  ")
	CheckErr(err)

	return string(data), nil
}
