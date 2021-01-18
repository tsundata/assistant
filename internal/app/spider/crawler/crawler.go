package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/cron"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/rule"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Crawler struct {
	jobs  map[string]rule.Rule
	outCh chan rule.Result

	rdb       *redis.Client
	logger    *zap.Logger
	msgClient pb.MessageClient
	midClient pb.MiddleClient
	subClient pb.SubscribeClient
}

func New(rdb *redis.Client, logger *zap.Logger,
	msgClient pb.MessageClient, midClient pb.MiddleClient, subClient pb.SubscribeClient) *Crawler {
	return &Crawler{
		jobs:      make(map[string]rule.Rule),
		outCh:     make(chan rule.Result),
		rdb:       rdb,
		logger:    logger,
		msgClient: msgClient,
		midClient: midClient,
		subClient: subClient,
	}
}

func (s *Crawler) LoadRule(p string) error {
	ctx := context.Background()
	return filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if ext := filepath.Ext(path); ext != ".yml" && ext != ".yaml" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var r rule.Rule
		err = yaml.Unmarshal(data, &r)
		if err != nil {
			return err
		}

		// check
		if r.Name == "" {
			return nil
		}
		if r.When == "" {
			return nil
		}
		if !utils.IsUrl(r.Page.URL) {
			return nil
		}

		// register
		_, err = s.subClient.Register(ctx, &pb.SubscribeRequest{
			Text: r.Name,
		})
		if err != nil {
			return err
		}
		s.jobs[r.Name] = r

		return nil
	})
}

func (s *Crawler) Daemon() {
	log.Println("subscribe spider cron starting...")

	for name, job := range s.jobs {
		log.Printf("spider %v: crawl...", name)
		go s.ruleWorker(name, job, s.outCh)
	}

	go s.resultWorker()
}

func (s *Crawler) resultWorker() {
	for out := range s.outCh {
		ctx := context.Background()
		latest := out.Result

		dataKey := fmt.Sprintf("%s:latest", out.Name)
		sendKey := fmt.Sprintf("%s:send", out.Name)

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
		for _, item := range out.Result {
			s.rdb.SAdd(ctx, dataKey, item)
		}
		s.rdb.Expire(ctx, dataKey, 7*24*time.Hour)

		// is instant
		if out.Instant {
			// send
			s.rdb.Set(ctx, sendKey, time.Now().Unix(), redis.KeepTTL)
			s.send(out.Name, diff)
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
			s.send(out.Name, diff)
		}
	}
}

func (s *Crawler) send(name string, out []string) {
	text := ""
	if len(out) <= 5 {
		text = fmt.Sprintf("Channel %s\n%s", name, strings.Join(out, "\n"))
	} else {
		// web page display
		j, err := json.Marshal(out)
		if err != nil {
			return
		}

		reply, err := s.midClient.CreatePage(context.Background(), &pb.PageRequest{
			Title:   fmt.Sprintf("Channel %s", name),
			Content: utils.ByteToString(j),
		})
		if err != nil {
			return
		}

		text = fmt.Sprintf("Channel %s\n%s\n %s", name, strings.Join(out[:5], "\n"), reply.GetText())
	}

	_, err := s.msgClient.Send(context.Background(), &pb.MessageRequest{
		Text: text,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Crawler) ruleWorker(name string, r rule.Rule, outCh chan rule.Result) {
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
			state, err := s.subClient.Status(context.Background(), &pb.SubscribeRequest{
				Text: name,
			})
			if err != nil {
				s.logger.Error(err.Error())
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
						log.Println("processSpiderRule panic", name, r)
					}
				}()
				return rule.RunRule(r)
			}()
			if len(result) > 0 {
				outCh <- rule.Result{
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
