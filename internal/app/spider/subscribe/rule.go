package subscribe

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tsundata/assistant/internal/app/spider/rule"
)

var Rules = map[string]rule.Rule{
	"news": {
		When:    "* * * * *",
		Instant: false,
		Action: func() []string {
			var result []string

			doc, err := rule.Document("https://www.v2ex.com/?tab=nodes")
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
		When:    "*/10 * * * *",
		Instant: false,
		Action: func() []string {
			return []string{}
		},
	},
}
