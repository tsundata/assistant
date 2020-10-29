package controllers

import (
	"context"
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
	gc.client.Dial(context.Background(), "Slack.SendMessage", "Hi Slack", &reply)

	c.JSON(http.StatusOK, framework.H{"time": time.Now().String(), "reply": reply})
}
