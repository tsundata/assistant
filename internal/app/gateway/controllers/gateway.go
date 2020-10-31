package controllers

import (
	"context"
	"fmt"
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
	var reply string
	_ = gc.client.Call(context.Background(), "Slack.SendMessage", "Hi Slack", &reply)
	fmt.Println(reply)

	var errCode int
	_ = gc.client.Call(context.Background(), "Subscribe.List", "", &errCode)
	fmt.Println(errCode)

	err := gc.client.Call(context.Background(), "Subscribe.Abc", "", &errCode)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, framework.H{"time": time.Now().String(), "reply": reply})
}
