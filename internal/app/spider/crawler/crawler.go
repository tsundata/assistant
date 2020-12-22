package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/rule"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.uber.org/zap"
	"log"
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

func (s *Crawler) Register(rules map[string]rule.Rule) {
	ctx := context.Background()
	for name, job := range rules {
		_, err := s.subClient.Register(ctx, &pb.SubscribeRequest{
			Text: name,
		})
		if err != nil {
			s.logger.Error(err.Error())
			continue
		}
		s.jobs[name] = job
	}
}

func (s *Crawler) Daemon() {
	log.Println("subscribe spider cron starting...")

	for name, job := range s.jobs {
		log.Printf("spider %v: crawl...", name)
		go rule.ProcessSpiderRule(name, job, s.outCh)
	}

	s.process()
}

func (s *Crawler) process() {
	go func() {
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
				s.Send(out.Name, diff)
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
				s.Send(out.Name, diff)
			}
		}
	}()
}

func (s *Crawler) Send(name string, out []string) {
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
			Content: string(j),
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
