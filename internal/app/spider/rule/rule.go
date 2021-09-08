package rule

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"sort"
	"strings"
)

type Rule struct {
	Name    string `yaml:"name"`
	Channel string `yaml:"channel"`
	When    string `yaml:"when"`
	Mode    string `yaml:"mode"`
	Page    struct {
		URL  string            `yaml:"url"`
		List string            `yaml:"list"`
		Item map[string]string `yaml:"item"`
	}
}

func (r Rule) Run() []string {
	var result []string

	doc, err := document(r.Page.URL)
	if err != nil {
		return result
	}

	doc.Find(r.Page.List).Each(func(i int, s *goquery.Selection) {
		// sort keys
		keys := make([]string, 0, len(r.Page.Item))
		for k := range r.Page.Item {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		txt := strings.Builder{}
		for _, k := range keys {
			f := ParseFun(s, r.Page.Item[k])
			v, err := f.Invoke()
			if err != nil {
				continue
			}
			v = strings.TrimSpace(v)
			v = strings.ReplaceAll(v, "\n", "")
			v = strings.ReplaceAll(v, "\r\n", "")
			if v == "" {
				continue
			}
			txt.WriteString(k)
			txt.WriteString(": ")
			txt.WriteString(v)
			txt.WriteString("\n")
		}
		if txt.Len() == 0 {
			return
		}
		result = append(result, txt.String())
	})
	return result
}

type Result struct {
	Name    string
	Channel string
	Mode    string
	Result  []string
}

func document(url string) (*goquery.Document, error) {
	res, err := http.Get(url) // #nosec
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}
