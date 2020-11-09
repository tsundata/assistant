package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"log"
	"net/http"
	"time"
)

type GatewayController struct {
	subClient *rpc.Client
	msgClient *rpc.Client
}

func NewGatewayController(subClient *rpc.Client, msgClient *rpc.Client) *GatewayController {
	return &GatewayController{subClient: subClient, msgClient: msgClient}
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
