package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
)

func CreateInitControllersFn(gc *GatewayController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/", gc.Index)
		r.GET("/foo", gc.Foo)
	}
}
