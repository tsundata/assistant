package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	slackVendor "github.com/tsundata/assistant/internal/pkg/vendors/slack"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GatewayController struct {
	o         *gateway.Options
	subClient *rpc.Client
	msgClient *rpc.Client
}

func NewGatewayController(o *gateway.Options, subClient *rpc.Client, msgClient *rpc.Client) *GatewayController {
	return &GatewayController{o: o, subClient: subClient, msgClient: msgClient}
}

func (gc *GatewayController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"path": "ROOT"})
}

func (gc *GatewayController) Foo(c *gin.Context) {
	args := &proto.Message{
		Id:    11,
		Input: "in ===>",
	}

	var reply proto.Message
	err := gc.subClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	log.Printf(reply.String())

	args = &proto.Message{
		Id:    11,
		Input: "in ===>",
	}

	err = gc.msgClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	log.Printf(reply.String())

	c.JSON(http.StatusOK, gin.H{"time": time.Now().String(), "reply": "reply"})
}

func (gc *GatewayController) SlackShortcut(c *gin.Context) {
	// verificationTokens
	s, err := slackVendor.SlashShortcutParse(c.Request)
	if err != nil {
		log.Println(err)
		return
	}

	if !s.ValidateToken(gc.o.Verification) {
		log.Println("unvalidated verificationTokens")
		return
	}

	if s.Type == "shortcut" {
		switch s.CallbackID {
		case "report":
			log.Println("report")
		}
	}

	if s.Type == "message_action" {
		switch s.CallbackID {
		case "delete":
			log.Println("delete")
		}
	}

	c.String(http.StatusOK, "OK")
	return
}

func (gc *GatewayController) SlackCommand(c *gin.Context) {
	// verificationTokens
	s, err := slack.SlashCommandParse(c.Request)
	if err != nil {
		log.Println(err)
		return
	}

	if !s.ValidateToken(gc.o.Verification) {
		log.Println("unvalidated verificationTokens")
		return
	}

	// parse
	switch s.Command {
	case "/view":
		id, err := strconv.Atoi(s.Text)
		if err != nil {
			log.Println(err)
			return
		}
		msg := &proto.Message{
			Id: uint64(id),
		}
		var reply proto.Message
		err = gc.msgClient.Call(context.Background(), "View", msg, &reply)
		if err != nil {
			log.Println(err)
			return
		}

		if reply.Id > 0 {
			err = slackVendor.ResponseText(s.ResponseURL, reply.Input)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = slackVendor.ResponseText(s.ResponseURL, "view failed")
			if err != nil {
				log.Println(err)
			}
		}
	}
	c.String(http.StatusOK, "OK")
	return
}

func (gc *GatewayController) SlackEvent(c *gin.Context) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	body := buf.String()

	api := slack.New(gc.o.Token)
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: gc.o.Verification}))
	if err != nil {
		log.Println(err)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"ok": false})
			return
		}
		c.String(http.StatusOK, r.Challenge)
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
					log.Println(err)
					return
				}
				if reply.Id > 0 {
					_, _, err = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("MGID: %d", reply.Id), false))
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// TODO
func (gc *GatewayController) AgentWebhook(c *gin.Context) {

}
