package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
)

func CreateInitControllersFn(gc *GatewayController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/", gc.Index)
		r.GET("/foo", gc.Foo)
		r.POST("/slack/shortcut", gc.SlackShortcut)
		r.POST("/slack/command", gc.SlackCommand)
		r.POST("/slack/event", gc.SlackEvent)
		r.POST("/agent/webhook", gc.AgentWebhook)
	}
}
