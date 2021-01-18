package rule

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Rule struct {
	Name    string `yaml:"name"`
	When    string `yaml:"when"`
	Instant bool   `yaml:"instant"`
	Page    struct {
		URL  string            `yaml:"url"`
		List string            `yaml:"list"`
		Item map[string]string `yaml:"item"`
	}
}

type Result struct {
	Name    string
	Instant bool
	Result  []string
}

func document(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func RunRule(r Rule) []string {
	var result []string

	doc, err := document(r.Page.URL)
	if err != nil {
		return result
	}

	doc.Find(r.Page.List).Each(func(i int, s *goquery.Selection) {
		txt := strings.Builder{}
		for k, v := range r.Page.Item {
			f := ParseFun(s, v)
			c, err := f.Invoke()
			if err != nil {
				log.Println(err)
				continue
			}
			txt.WriteString(k)
			txt.WriteString(": ")
			txt.WriteString(c)
			txt.WriteString("\n")
		}
		result = append(result, txt.String())
	})
	return result
}
