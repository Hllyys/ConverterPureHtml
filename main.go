package main

import (
	"PureHtmlConverter/scraper"
	"fmt"
)

func main() {
	html :=

		`
<div class="product">
  <h2 class="title">Hello World</h2>
  <span class="price">₺99</span>

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
		//   <h2 class="title">Hello World</h2>
		//   <span class="price">₺1</span>
		//   <div class="details">
		//     <span class="date">2025-05-01</span>
		//     <div class="seller">
		//       <span class="name">John Doe</span>
		//       <span class="rating">4.5</span>
		//     </div>
		//   </div>
		//   <div class="tags">
		//     <span>electronics</span>
		//     <span>gadgets</span>
		//     <span>sale</span>
		//   </div>
		// </div>
	yamlStr := `
selector: ".product"
fields:
  title:
    type: primitive
    selector: "h2"
    transform: trim

  price:
    type: union
    union:
      - type: primitive
        selector: ".price"
        transform: trim
      - type: constant
        constant: "₺0"

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

  tags:
    type: array
    selector: ".tags span"
    item:
      type: primitive
      selector: ""
      transform: trim

`
	//selector: ".product"
	// fields:
	//   title:
	//     type: primitive
	//     selector: "h2"
	//     transform: trim

	//   source:
	//     type: constant
	//     constant: "example.com"

	//   details:
	//     type: object
	//     selector: ".details"
	//     fields:
	//       date:
	//         type: primitive
	//         selector: ".date"
	//         transform: ["trim", "date"]

	//       seller:
	//         type: object
	//         selector: ".seller"
	//         fields:
	//           name:
	//             type: primitive
	//             selector: ".name"
	//             transform: trim
	//           rating:
	//             type: primitive
	//             selector: ".rating"
	//             transform: trim

	//   tags:
	//     type: array
	//     selector: ".tags span"
	//     item:
	//       type: primitive
	//       selector: ""
	//       transform: trim

	//   price:
	//     type: union
	//     union:
	//       - type: primitive
	//         selector: ".price"
	//         transform: trim
	//       - type: primitive
	//         selector: ".alt-price"
	//         transform: trim
	//       - type: constant
	//         constant: "₺0"
	config, err := scraper.ParseDynamicYAML(yamlStr)
	scraper.CheckFatal(err, "YAML parse hatası")

	result, err := scraper.ExtractWithConfig(html, config, "https://example.com/page.html")
	scraper.CheckFatal(err, "Extract hatası")

	fmt.Println(result)
}
