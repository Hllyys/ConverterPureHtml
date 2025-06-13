package main

import (
	"Converter/scraper"
	"fmt"
)

func main() {
	html := `
<div class="product">
  <h2 class="title">Hello World</h2>
  <span class="price">₺1</span>
  <div class="details">
    <span class="date">2025-05-01</span>
    <div class="seller">
      <span class="name">John Doe</span>
      <span class="rating">4.5</span>
    </div>
  </div>
  <div class="tags">
    <span>electronics</span>
    <span>gadgets</span>
    <span>sale</span>
  </div>
</div>

`

	yaml := `
selector: ".product"
fields:

  # primitive örneği
  title:
    type: primitive
    selector: "h2"
    transform: trim

  # constant örneği
  source:
    type: constant
    constant: "example.com"

  # object örneği
  details:
    type: object
    selector: ".details"
    fields:
      date:
        type: primitive
        selector: ".date"
        transform: ["trim", "date"]

      seller:
        type: object
        selector: ".seller"
        fields:
          name:
            type: primitive
            selector: ".name"
            transform: trim
          rating:
            type: primitive
            selector: ".rating"
            transform: trim

  # array örneği
  tags:
    type: array
    selector: ".tags span"
    item:
      type: primitive
      selector: ""
      transform: trim

  # union örneği
  price:
    type: union
    union:
      - type: primitive
        selector: ".price"
        transform: trim
      - type: primitive
        selector: ".alt-price"
        transform: trim
      - type: constant
        constant: "₺0"  

`

	result, err := scraper.Extract(html, yaml, "https://example.com/page.html")
	scraper.CheckFatal(err, "Extract işleminde hata oluştu.")
	fmt.Println(result)
}
