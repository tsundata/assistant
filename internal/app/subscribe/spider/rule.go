package spider

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
)

var SubscribeRules = map[string]Rule{
	"news": {
		false,
		"* * * * *",
		func() []string {
			var result []string
			log.Println("demo 1 minute spider " + time.Now().String())

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
		"*/3 * * * *",
		func() []string {
			log.Println("demo 3 minute spider " + time.Now().String())
			return []string{}
		},
	},
}
