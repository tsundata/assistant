package spider

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v8"
	"github.com/gorhill/cronexpr"
	"github.com/spf13/viper"
	httpPkg "github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Spider struct {
	rdb   *redis.Client
	outCh chan Result

	webhook string
}

func New(rdb *redis.Client, v *viper.Viper) *Spider {
	slack := v.GetStringMapString("slack")
	return &Spider{
		rdb:     rdb,
		outCh:   make(chan Result),
		webhook: slack["webhook"],
	}
}

func (s *Spider) Cron() {
	log.Println("subscribe spider cron starting...")

	for name, rule := range spiderRules {
		log.Printf("spider %v: crawl...", name)
		go processSpiderRule(name, rule, s.outCh)
	}

	s.process()
}

func (s *Spider) process() {
	go func() {
		for {
			select {
			case out := <-s.outCh:
				fmt.Println(out)
				ctx := context.Background()
				latest := out.result

				dataKey := fmt.Sprintf("%s:latest", out.name)
				sendKey := fmt.Sprintf("%s:send", out.name)

				smembers := s.rdb.SMembers(ctx, dataKey)
				old, err := smembers.Result()
				if err != nil {
					continue
				}

				// diff
				diff := utils.SliceDiff(latest, old)
				if len(diff) == 0 {
					continue
				}
				// add data
				for _, item := range out.result {
					s.rdb.SAdd(ctx, dataKey, item)
				}
				// FIXME
				s.rdb.Expire(ctx, dataKey, 7*24*time.Hour)

				// is instant
				if out.instant {
					// send
					s.rdb.Set(ctx, sendKey, time.Now().Unix(), redis.KeepTTL)
					s.Send(out.name, diff)
				} else {
					sendStringCmd := s.rdb.Get(ctx, sendKey)
					sendString, err := sendStringCmd.Result()
					if err != nil {
						continue
					}
					oldSend, err := strconv.ParseInt(sendString, 10, 64)
					// FIXME
					fmt.Println(oldSend)
					fmt.Println(time.Now().Unix())
					if err != nil {
						continue
					}

					if time.Now().Unix()-oldSend < 24*3600 {
						continue
					}

					s.rdb.Set(ctx, sendKey, time.Now().Unix(), redis.KeepTTL)
					s.Send(out.name, diff)
				}
			default:

			}
		}
	}()
}

func (s *Spider) Send(name string, out []string) {
	fmt.Printf("send event : %v\n", out)

	text := ""
	if len(out) <= 5 {
		text = fmt.Sprintf("Channel %s\n%s", name, strings.Join(out, "\n"))
	} else {
		// TODO web display
		text = fmt.Sprintf("Channel %s\n%s\n http://demo.url/abc", name, strings.Join(out[:5], "\n"))
	}

	client := httpPkg.NewClient()
	resp, err := client.PostJSON(s.webhook, map[string]interface{}{
		"text": text,
	})

	if err != nil {
		log.Println(err)
		return
	}

	fasthttp.ReleaseResponse(resp)
}

type Result struct {
	name    string
	instant bool
	result  []string
}

func processSpiderRule(name string, rule Rule, outCh chan Result) {
	nextTime := cronexpr.MustParse(rule.When).Next(time.Now())
	for {
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			result := rule.Action()
			outCh <- Result{
				name:    name,
				instant: rule.Instant,
				result:  result,
			}
		}
		nextTime = cronexpr.MustParse(rule.When).Next(time.Now())
		time.Sleep(2 * time.Second)
	}
}

type Rule struct {
	Instant bool
	When    string
	Action  func() []string
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
