package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

var spiderRules = map[string]Rule{
	"demo 1 minute": {
		false,
		"* * * * *",
		func() []string {
			var result []string
			fmt.Println("demo 1 minute spider " + time.Now().String())

			doc, err := document("https://www.v2ex.com/")
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
	"demo 3 minute": {
		false,
		"*/3 * * * *",
		func() []string {
			fmt.Println("demo 3 minute spider " + time.Now().String())
			return []string{}
		},
	},
}
