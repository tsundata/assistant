package spider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v8"
	"github.com/gorhill/cronexpr"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Spider struct {
	rdb   *redis.Client
	outCh chan Result

	msgClient *pb.MessageClient
	midClient *pb.MiddleClient
}

func New(rdb *redis.Client, msgClient *pb.MessageClient, midClient *pb.MiddleClient) *Spider {
	return &Spider{
		rdb:       rdb,
		outCh:     make(chan Result),
		msgClient: msgClient,
		midClient: midClient,
	}
}

func (s *Spider) Cron() {
	log.Println("subscribe spider cron starting...")

	for name, rule := range SubscribeRules {
		log.Printf("spider %v: crawl...", name)
		go processSpiderRule(name, rule, s.outCh)
	}

	s.process()
}

func (s *Spider) process() {
	go func() {
		for out := range s.outCh {
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
			diff := utils.SliceDiff(old, latest)
			if len(old) > 0 && len(diff) == 0 {
				continue
			} else {
				diff = latest
			}
			if len(diff) == 0 {
				continue
			}

			// add data
			for _, item := range out.result {
				s.rdb.SAdd(ctx, dataKey, item)
			}
			s.rdb.Expire(ctx, dataKey, 7*24*time.Hour)

			// is instant
			if out.instant {
				// send
				s.rdb.Set(ctx, sendKey, time.Now().Unix(), redis.KeepTTL)
				s.Send(out.name, diff)
			} else {
				sendStringCmd := s.rdb.Get(ctx, sendKey)
				sendString, _ := sendStringCmd.Result()
				oldSend := int64(0)
				if sendString != "" {
					oldSend, _ = strconv.ParseInt(sendString, 10, 64)
				}

				if time.Now().Unix()-oldSend < 24*3600 {
					continue
				}

				s.rdb.Set(ctx, sendKey, time.Now().Unix(), redis.KeepTTL)
				s.Send(out.name, diff)
			}
		}
	}()
}

func (s *Spider) Send(name string, out []string) {
	text := ""
	if len(out) <= 5 {
		text = fmt.Sprintf("Channel %s\n%s", name, strings.Join(out, "\n"))
	} else {
		// web page display
		j, err := json.Marshal(out)
		if err != nil {
			return
		}

		reply, err := (*s.midClient).CreatePage(context.Background(), &pb.PageRequest{
			Title:   fmt.Sprintf("Channel %s", name),
			Content: string(j),
		})
		if err != nil {
			return
		}

		text = fmt.Sprintf("Channel %s\n%s\n %s", name, strings.Join(out[:5], "\n"), reply.GetText())
	}

	_, err := (*s.msgClient).Send(context.Background(), &pb.MessageRequest{
		Text: text,
	})
	if err != nil {
		log.Println(err)
		return
	}
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
