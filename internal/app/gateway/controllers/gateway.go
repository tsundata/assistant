package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/tsundata/assistant/api"
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
	var reply int
	gc.client.Dial(context.Background(), "Foo.Sum", &api.Args{
		Num1: 10,
		Num2: 10,
	}, &reply)

	c.JSON(http.StatusOK, framework.H{"foo": time.Now().String(), "reply": reply})
}
