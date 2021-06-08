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
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors/telegram"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type GatewayController struct {
	opt    *config.AppConfig
	rdb    *redis.Client
	logger *logger.Logger
	client *rpc.Client
}

func NewGatewayController(opt *config.AppConfig, rdb *redis.Client, logger *logger.Logger, client *rpc.Client) *GatewayController {
	return &GatewayController{
		opt:    opt,
		rdb:    rdb,
		logger: logger,
		client: client,
	}
}

func (gc *GatewayController) Index(c *fiber.Ctx) error {
	return c.SendString("Gateway")
}

func (gc *GatewayController) SlackEvent(c *fiber.Ctx) error {
	body := c.Request().Body()

	api := slack.New(gc.opt.Slack.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(body, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.opt.Slack.Verification}))
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			gc.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString(err.Error())
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
				rKey := "message:repeated:slack:" + ev.ClientMsgID
				isRepeated := gc.rdb.Get(context.Background(), rKey).Val()
				if len(isRepeated) > 0 {
					return c.Status(http.StatusBadRequest).SendString("repeat message")
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

				reply, err := rpcclient.GetMessageClient(gc.client).Create(context.Background(), &pb.MessageRequest{
					Uuid: ev.ClientMsgID,
					Text: ev.Text,
				})
				if err != nil {
					gc.logger.Error(err)
					return c.Status(http.StatusBadRequest).SendString(err.Error())
				}

				if reply.GetId() > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("ID: %d", reply.GetId()), false))
				} else {
					for _, item := range reply.GetText() {
						if item != "" {
							_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(item, false))
						}
					}
				}
				if err != nil {
					gc.logger.Error(err)
					return c.Status(http.StatusBadRequest).SendString(err.Error())
				}
			}
		}
	}

	return c.SendString("OK")
}

func (gc *GatewayController) TelegramEvent(c *fiber.Ctx) error {
	var incoming telegram.IncomingRequest
	err := json.Unmarshal(c.Request().Body(), &incoming)
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// todo auth

	// ignore repeated message
	rKey := fmt.Sprintf("message:repeated:telegram:%d", incoming.UpdateId)
	isRepeated := gc.rdb.Get(context.Background(), rKey).Val()
	if len(isRepeated) > 0 {
		return c.Status(http.StatusBadRequest).SendString("repeat message")
	}
	gc.rdb.Set(context.Background(), rKey, time.Now().Unix(), 7*24*time.Hour)

	// empty message
	if incoming.Message.Text == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	// handle message
	uuid, err := utils.GenerateUUID()
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	reply, err := rpcclient.GetMessageClient(gc.client).Create(context.Background(), &pb.MessageRequest{
		Uuid: uuid,
		Text: incoming.Message.Text,
	})
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// reply message
	api := telegram.NewTelegram(gc.opt.Telegram.Token)
	if reply.GetId() > 0 {
		_, err = api.SendMessage(incoming.Message.Chat.Id, fmt.Sprintf("ID: %d", reply.GetId()))
	} else {
		for _, item := range reply.GetText() {
			if item != "" {
				_, err = api.SendMessage(incoming.Message.Chat.Id, item)
			}
		}
	}
	if err != nil {
		gc.logger.Error(err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

func (gc *GatewayController) GetPage(c *fiber.Ctx) error {
	var in pb.PageRequest
	err := c.QueryParser(&in)
	if err != nil {
		return err
	}

	reply, err := rpcclient.GetMiddleClient(gc.client).GetPage(context.Background(), &in)
	if err != nil {
		return err
	}
	return c.JSON(reply)
}
