package controllers

import (
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/framework"
)

func CreateInitControllersFn(gc *GatewayController) http.InitControllers {
	return func(r *framework.Engine) {
		r.GET("/", gc.Index)
		r.GET("/foo", gc.Foo)
	}
}
