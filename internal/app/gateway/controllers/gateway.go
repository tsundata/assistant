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
	"github.com/tsundata/assistant/internal/app/gateway/scheduler"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	slackVendor "github.com/tsundata/assistant/internal/pkg/vendors/slack"
	"log"
	"net/http"
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
	args := &proto.Detail{
		Id:          11,
		Name:        "in ===>",
		Price:       29292,
		CreatedTime: nil,
	}

	var reply proto.Detail
	err := gc.subClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	log.Printf(reply.String())

	args = &proto.Detail{
		Id:          11,
		Name:        "in ===>",
		Price:       29292,
		CreatedTime: nil,
	}

	err = gc.msgClient.Call(context.Background(), "Open", args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	log.Printf(reply.String())

	c.JSON(http.StatusOK, gin.H{"time": time.Now().String(), "reply": "reply"})
}

// TODO
func (gc *GatewayController) SlackShortcut(c *gin.Context) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	body := buf.String()

	fmt.Println(gc.o.Signing)
	err = slackVendor.SecretsVerifier(c.Request.Header, buf.Bytes(), gc.o.Signing)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	fmt.Println(c.Request.Header)
	fmt.Println(body)
}

// TODO
func (gc *GatewayController) SlackCommand(c *gin.Context) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}
	body := buf.String()

	fmt.Println(gc.o.Signing)
	err = slackVendor.SecretsVerifier(c.Request.Header, buf.Bytes(), gc.o.Signing)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"ok": false})
		return
	}

	fmt.Println(c.Request.Header)
	fmt.Println(body)
}

// TODO
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

	fmt.Println(eventsAPIEvent)
	fmt.Println(eventsAPIEvent.InnerEvent.Type)
	fmt.Println(eventsAPIEvent.InnerEvent.Data)

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
		case *slackevents.AppMentionEvent:
			api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		case *slackevents.MessageEvent:
			fmt.Println(ev.Text)
			if ev.BotID == "" {
				scheduler.EventScheduler(ev.Text)

				api.PostMessage(ev.Channel, slack.MsgOptionText("Text is "+ev.Text, false))
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// TODO
func (gc *GatewayController) AgentWebhook(c *gin.Context) {

}
