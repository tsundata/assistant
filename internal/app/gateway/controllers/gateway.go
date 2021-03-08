package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type GatewayController struct {
	opt       *gateway.Options
	rdb       *redis.Client
	logger    *zap.Logger
	subClient pb.SubscribeClient
	msgClient pb.MessageClient
}

func NewGatewayController(opt *gateway.Options, rdb *redis.Client, logger *zap.Logger, subClient pb.SubscribeClient, msgClient pb.MessageClient) *GatewayController {
	return &GatewayController{opt: opt, rdb: rdb, logger: logger, subClient: subClient, msgClient: msgClient}
}

func (gc *GatewayController) Index(c *fiber.Ctx) error {
	return c.SendString("Gateway")
}

func (gc *GatewayController) SlackEvent(c *fiber.Ctx) error {
	body := c.Request().Body()

	api := slack.New(gc.opt.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(body, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.opt.Verification}))
	if err != nil {
		gc.logger.Error(err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			gc.logger.Error(err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}
		return c.SendString(r.Challenge)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			// ignore bot message
			if ev.ClientMsgID != "" {
				// ignore repeated message
				rKey := "message:repeated:" + ev.ClientMsgID
				isRepeated := gc.rdb.Get(context.Background(), rKey).Val()
				if len(isRepeated) > 0 {
					return c.SendStatus(http.StatusBadRequest)
				}
				gc.rdb.Set(context.Background(), rKey, time.Now().Unix(), 7*24*time.Hour)

				// special <url>, utf8 whitespace
				re := regexp.MustCompile("<" + utils.UrlRegex + ">")
				urls := re.FindAllString(ev.Text, -1)
				for _, url := range urls {
					ev.Text = strings.ReplaceAll(ev.Text, url, strings.TrimRight(strings.TrimLeft(url, "<"), ">"))
				}
				re = regexp.MustCompile(`[\s\p{Zs}]+`)
				ev.Text = re.ReplaceAllString(ev.Text, " ")

				reply, err := gc.msgClient.Create(context.Background(), &pb.MessageRequest{
					Uuid: ev.ClientMsgID,
					Text: ev.Text,
				})
				if err != nil {
					gc.logger.Error(err.Error())
					return c.SendStatus(http.StatusBadRequest)
				}

				if reply.GetId() > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("ID: %d", reply.GetId()), false))
				} else {
					for _, item := range reply.GetText() {
						_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(item, false))
					}
				}
				if err != nil {
					gc.logger.Error(err.Error())
					return c.SendStatus(http.StatusBadRequest)
				}
			}
		}
	}

	return c.SendString("OK")
}
