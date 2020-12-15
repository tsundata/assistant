package spider

import (
	"github.com/PuerkitoBio/goquery"
)

var SubscribeRules = map[string]Rule{
	"news": {
		false,
		"* * * * *",
		func() []string {
			var result []string

			doc, err := document("https://www.v2ex.com/?tab=nodes")
			if err != nil {
				return result
			}

			doc.Find(".box .item").Each(func(i int, s *goquery.Selection) {
				title := s.Find(".item_title a").Text()
				result = append(result, title)
			})
			return result
		},
	},
	"demo": {
		false,
		"*/10 * * * *",
		func() []string {
			return []string{}
		},
	},
}
