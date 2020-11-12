package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/app/gateway"
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
	panic("222")
	args := &proto.Message{
		Id:    11,
		Input: "in ===>",
	}

	var reply proto.Message
	err := gc.subClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		gc.logger.Error(err.Error())
	}

	gc.logger.Info(reply.String())

	args = &proto.Message{
		Id:    11,
		Input: "in ===>",
	}

	err = gc.msgClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		gc.logger.Error(err.Error())
	}

	gc.logger.Info(reply.String())

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
		msg := &proto.Message{
			Id: uint64(id),
		}
		var reply proto.Message
		err = gc.msgClient.Call(context.Background(), "View", msg, &reply)
		if err != nil {
			gc.logger.Error(err.Error())
			c.Error(err.Error(), http.StatusBadRequest)
			return
		}

		if reply.Id > 0 {
			err = slackVendor.ResponseText(s.ResponseURL, reply.Input)
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
	}

	c.Response.SetBodyString("OK")
}

func (gc *GatewayController) SlackEvent(c *fasthttp.RequestCtx) {
	body := c.Request.Body()

	api := slack.New(gc.o.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.o.Verification}))
	if err != nil {
		gc.logger.Error(err.Error())
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
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
				msg := &proto.Message{
					Id:        0,
					Uuid:      ev.ClientMsgID,
					ChannelId: ev.Channel,
					Input:     ev.Text,
				}
				var reply proto.Message
				err = gc.msgClient.Call(context.Background(), "Create", msg, &reply)
				if err != nil {
					gc.logger.Error(err.Error())
					return
				}
				if reply.Id > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("MGID: %d", reply.Id), false))
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
