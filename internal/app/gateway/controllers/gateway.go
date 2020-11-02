package controllers

import (
	"context"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/framework"
)

type GatewayController struct {
	client *rpc.Client
}

func NewGatewayController(client *rpc.Client) *GatewayController {
	return &GatewayController{client: client}
}

func (gc *GatewayController) Index(c *framework.Context) {
	c.JSON(http.StatusOK, framework.H{"path": "ROOT"})
}

func (gc *GatewayController) Foo(c *framework.Context) {
	//var reply string
	//_ = gc.client.Call(context.Background(), "Slack.SendMessage", "Hi Slack", &reply)
	//log.Println(reply)
	//
	//var errCode int
	//_ = gc.client.Call(context.Background(), "Subscribe.List", "", &errCode)
	//log.Println(errCode)

	payload := &proto.Detail{
		Id:          2828,
		Name:        "=====> in",
		Price:       212211,
		CreatedTime: nil,
	}
	args, _ := utils.ProtoMarshal(payload)

	var replay *[]byte
	err := gc.client.Call(context.Background(), "Subscribe.Subscribe.Open", args, &replay)
	if err != nil {
		log.Println(err)
	}
	if replay == nil {
		log.Println("error replay")
	}

	var d proto.Detail
	err = utils.ProtoUnmarshal(*replay, &d)
	if err != nil {
		log.Println(err)
	}

	log.Println(d)

	c.JSON(http.StatusOK, framework.H{"time": time.Now().String(), "reply": "reply"})
}
