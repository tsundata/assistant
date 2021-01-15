package rule

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/influxdata/cron"
	"log"
	"net/http"
	"strings"
	"time"
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

func runRule(r Rule) []string {
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

type Result struct {
	Name    string
	Instant bool
	Result  []string
}

func ProcessSpiderRule(name string, r Rule, outCh chan Result) {
	p, err := cron.ParseUTC(r.When)
	if err != nil {
		log.Println(err)
		return
	}
	nextTime, err := p.Next(time.Now())
	if err != nil {
		log.Println(err)
		return
	}
	for {
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			result := func() []string {
				defer func() {
					if r := recover(); r != nil {
						log.Println("processSpiderRule panic", name, r)
					}
				}()
				return runRule(r)
			}()
			if len(result) > 0 {
				outCh <- Result{
					Name:    name,
					Instant: r.Instant,
					Result:  result,
				}
			}
		}
		nextTime, err = p.Next(time.Now())
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}
