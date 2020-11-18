package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	slackVendor "github.com/tsundata/assistant/internal/pkg/vendors/slack"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type GatewayController struct {
	o         *gateway.Options
	logger    *zap.Logger
	subClient *rpc.Client
	msgClient *rpc.Client
}

func NewGatewayController(o *gateway.Options, logger *zap.Logger, subClient *rpc.Client, msgClient *rpc.Client) *GatewayController {
	return &GatewayController{o: o, logger: logger, subClient: subClient, msgClient: msgClient}
}

func (gc *GatewayController) Index(c *fasthttp.RequestCtx) {
	c.Response.SetBody([]byte("ROOT"))
}

func (gc *GatewayController) Foo(c *fasthttp.RequestCtx) {
	args := &model.Message{
		Content: "input --->",
	}

	var reply model.Message
	err := gc.subClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		gc.logger.Error(err.Error())
	}

	gc.logger.Info(reply.Content)

	c.Response.SetBodyString(time.Now().String())
}

func (gc *GatewayController) SlackShortcut(c *fasthttp.RequestCtx) {
	// verificationTokens
	s, err := slackVendor.SlashShortcutParse(&c.Request)
	if err != nil {
		gc.logger.Error(err.Error())
		return
	}

	if !s.ValidateToken(gc.o.Verification) {
		gc.logger.Info("unvalidated verificationTokens")
		return
	}

	if s.Type == "shortcut" {
		switch s.CallbackID {
		case "report":
			gc.logger.Info("report")
		}
	}

	if s.Type == "message_action" {
		switch s.CallbackID {
		case "delete":
			gc.logger.Info("delete")
		case "run":
			// TODO
			fmt.Println(s)
			gc.logger.Info("run")
			var reply string
			err := gc.msgClient.Call(context.Background(), "Run", s.Message.ClientMsgID, &reply)
			if err != nil {
				gc.logger.Error(err.Error())
				return
			}
			err = slackVendor.ResponseText(s.ResponseURL, reply)
			if err != nil {
				gc.logger.Error(err.Error())
				return
			}
		}
	}

	c.Response.SetBodyString("OK")
}

func (gc *GatewayController) SlackCommand(c *fasthttp.RequestCtx) {
	// verificationTokens
	s, err := slackVendor.SlashCommandParse(&c.Request)
	if err != nil {
		gc.logger.Error(err.Error())
		c.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if !s.ValidateToken(gc.o.Verification) {
		gc.logger.Info("unvalidated verificationTokens")
		c.Error("unvalidated verificationTokens", http.StatusBadRequest)
		return
	}

	// parse
	switch s.Command {
	case "/view":
		id, err := strconv.Atoi(s.Text)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}
		msg := &model.Message{
			ID: id,
		}
		var reply model.Message
		err = gc.msgClient.Call(context.Background(), "View", msg, &reply)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		if reply.ID > 0 {
			err = slackVendor.ResponseText(s.ResponseURL, reply.Content)
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			err = slackVendor.ResponseText(s.ResponseURL, "view failed")
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		}
	case "/run":
		// TODO
		id, err := strconv.Atoi(s.Text)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}
		msg := &model.Message{
			ID: id,
		}
		var reply model.Message
		err = gc.msgClient.Call(context.Background(), "View", msg, &reply)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		if reply.ID > 0 {
			var r string
			err = gc.msgClient.Call(context.Background(), "Run", reply.UUID, &r)
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
			err = slackVendor.ResponseText(s.ResponseURL, r)
			if err != nil {
				gc.logger.Error(err.Error())
				c.Error(err.Error(), http.StatusBadRequest)
				return
			}
		}
	}

	c.Response.SetBodyString("OK")
}

func (gc *GatewayController) SlackEvent(c *fasthttp.RequestCtx) {
	body := c.Request.Body()

	api := slack.New(gc.o.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(body, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.o.Verification}))
	if err != nil {
		gc.logger.Error(err.Error())
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}
		c.Response.SetBodyString(r.Challenge)
		return
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			// ignore bot message
			if ev.ClientMsgID != "" {
				msg := &model.Message{
					ID:          0,
					UUID:        ev.ClientMsgID,
					ChannelID:   ev.Channel,
					ChannelName: ev.ChannelType,
					Content:     ev.Text,
				}
				var reply model.Message
				err = gc.msgClient.Call(context.Background(), "Create", msg, &reply)
				if err != nil {
					gc.logger.Error(err.Error())
					return
				}
				if reply.ID > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("MGID: %d", reply.ID), false))
					if err != nil {
						gc.logger.Error(err.Error())
						return
					}
				}
			}
		}
	}

	c.Response.SetBodyString("OK")
}

// TODO
func (gc *GatewayController) AgentWebhook(c *fasthttp.RequestCtx) {

}
