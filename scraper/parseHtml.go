package scraper

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func ParseHTMLFromString(html string) (*goquery.Document, error) {
	reader := strings.NewReader(html)
	return goquery.NewDocumentFromReader(reader)
}
