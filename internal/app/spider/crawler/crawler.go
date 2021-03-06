package crawler

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/rule"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
	"time"
)

type Crawler struct {
	outCh chan rule.Result
	jobs  map[string]rule.Rule

	c         *config.AppConfig
	rdb       *redis.Client
	logger    *logger.Logger
	subscribe pb.SubscribeClient
	middle    pb.MiddleClient
	message   pb.MessageClient
}

func New() *Crawler {
	return &Crawler{
		jobs:  make(map[string]rule.Rule),
		outCh: make(chan rule.Result, 10),
	}
}

func (s *Crawler) SetService(
	c *config.AppConfig,
	rdb *redis.Client,
	logger *logger.Logger,
	subscribe pb.SubscribeClient,
	middle pb.MiddleClient,
	message pb.MessageClient) {
	s.c = c
	s.rdb = rdb
	s.logger = logger
	s.subscribe = subscribe
	s.middle = middle
	s.message = message
}

func (s *Crawler) LoadRule() error {
	data, err := s.c.GetConfig(fmt.Sprintf("%s/rules", app.Spider))
	if err != nil {
		return err
	}

	ruleYMLs := strings.Split(data, "---")

	ctx := context.Background()
	for _, ruleYML := range ruleYMLs {
		var r rule.Rule
		err = yaml.Unmarshal(util.StringToByte(ruleYML), &r)
		if err != nil {
			s.logger.Error(err)
			continue
		}

		// check
		if r.Name == "" {
			continue
		}
		if r.When == "" {
			continue
		}
		if !util.IsUrl(r.Page.URL) {
			continue
		}

		// register
		_, err = s.subscribe.Register(ctx, &pb.SubscribeRequest{
			Text: r.Name,
		})
		if err != nil {
			s.logger.Error(err)
			continue
		}
		s.jobs[r.Name] = r
	}
	return nil
}

func (s *Crawler) Daemon() {
	s.logger.Info("subscribe spider starting...")

	for name, job := range s.jobs {
		s.logger.Info("spider " + name + ": crawl...")
		go s.ruleWorker(name, job)
	}

	go s.resultWorker()
}

func (s *Crawler) ruleWorker(name string, r rule.Rule) {
	p, err := cron.ParseUTC(r.When)
	if err != nil {
		s.logger.Error(err)
		return
	}
	nextTime, err := p.Next(time.Now())
	if err != nil {
		s.logger.Error(err)
		return
	}
	for {
		if nextTime.Format("2006-01-02 15:04") == time.Now().Format("2006-01-02 15:04") {
			state, err := s.subscribe.Status(context.Background(), &pb.SubscribeRequest{
				Text: name,
			})
			if err != nil {
				s.logger.Error(err)
				continue
			}
			// unsubscribe
			if !state.State {
				time.Sleep(30 * time.Second)
				continue
			}

			result := func() []string {
				defer func() {
					if r := recover(); r != nil {
						s.logger.Warn("ruleWorker recover " + name)
						if v, ok := r.(error); ok {
							s.logger.Error(v)
						}
					}
				}()
				return r.Run()
			}()
			if len(result) > 0 {
				s.outCh <- rule.Result{
					Name:    name,
					Channel: r.Channel,
					Mode:    r.Mode,
					Result:  result,
				}
			}
		}
		nextTime, err = p.Next(time.Now())
		if err != nil {
			s.logger.Error(err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func (s *Crawler) resultWorker() {
	for out := range s.outCh {
		// filter
		diff := s.filter(out.Name, out.Mode, out.Result)
		// send
		s.send(out.Channel, out.Name, diff)
	}
}

func (s *Crawler) filter(name, mode string, latest []string) []string {
	ctx := context.Background()
	sentKey := fmt.Sprintf("spider:%s:sent", name)
	todoKey := fmt.Sprintf("spider:%s:todo", name)
	sendTimeKey := fmt.Sprintf("spider:%s:sendtime", name)

	// sent
	smembers := s.rdb.SMembers(ctx, sentKey)
	old, err := smembers.Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		s.logger.Error(err)
		return []string{}
	}

	// to do
	smembers = s.rdb.SMembers(ctx, todoKey)
	todo, err := smembers.Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		s.logger.Error(err)
		return []string{}
	}

	// merge
	tobeCompared := append(old, todo...)

	// diff
	diff := util.StringSliceDiff(latest, tobeCompared)

	switch mode {
	case "instant":
		s.rdb.Set(ctx, sendTimeKey, time.Now().Unix(), 0)
	case "daily":
		sendString := s.rdb.Get(ctx, sendTimeKey).Val()
		oldSend := int64(0)
		if sendString != "" {
			oldSend, _ = strconv.ParseInt(sendString, 10, 64)
		}

		if time.Now().Unix()-oldSend < 24*60*60 {
			for _, item := range diff {
				s.rdb.SAdd(ctx, todoKey, item)
			}

			return []string{}
		}

		diff = append(diff, todo...)

		s.rdb.Set(ctx, sendTimeKey, time.Now().Unix(), 0)
	default:
		return []string{}
	}

	// add data
	for _, item := range diff {
		s.rdb.SAdd(ctx, sentKey, item)
	}
	s.rdb.Expire(ctx, sentKey, 7*24*time.Hour)

	// clear to do
	s.rdb.Del(ctx, todoKey)

	return diff
}

func (s *Crawler) send(channel, name string, out []string) {
	if len(out) == 0 {
		return
	}

	// check send
	key := fmt.Sprintf("spider:send:%x", md5.Sum(util.StringToByte(strings.Join(out, "\n"))))
	isSet, err := s.rdb.SetNX(context.Background(), key, time.Now().Unix(), 24*time.Hour).Result()
	if err != nil || !isSet {
		return
	}

	// simplify
	text := ""
	if len(out) <= 5 {
		text = fmt.Sprintf("Channel %s (v%s)\n%s", name, version.Version, strings.Join(out, "\n"))
	} else {
		// web page display
		j, err := json.Marshal(out)
		if err != nil {
			return
		}

		reply, err := s.middle.CreatePage(context.Background(), &pb.PageRequest{
			Type:    "json",
			Title:   fmt.Sprintf("Channel %s (%s)", name, time.Now().Format("2006-01-02 15:04:05")),
			Content: util.ByteToString(j),
		})
		if err != nil {
			return
		}

		text = fmt.Sprintf("Channel %s (v%s)\n%s\n %s", name, version.Version, strings.Join(out[:5], "\n"), reply.GetText())
	}

	// send
	_, err = s.message.Send(context.Background(), &pb.MessageRequest{
		Channel: channel,
		Text:    text,
	})
	if err != nil {
		s.logger.Error(err)
		return
	}
}
