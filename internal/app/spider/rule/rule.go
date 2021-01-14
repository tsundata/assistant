package rule

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/influxdata/cron"
	"log"
	"net/http"
	"time"
)

type Rule struct {
	When    string
	Instant bool
	Action  func() []string
}

type Result struct {
	Name    string
	Instant bool
	Result  []string
}

func ProcessSpiderRule(name string, rule Rule, outCh chan Result) {
	p, err := cron.ParseUTC(rule.When)
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
				return rule.Action()
			}()
			if len(result) > 0 {
				outCh <- Result{
					Name:    name,
					Instant: rule.Instant,
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

func Document(url string) (*goquery.Document, error) {
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
