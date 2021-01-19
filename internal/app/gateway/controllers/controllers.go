package controllers

import (
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"log"
)

func CreateInitControllersFn(gc *GatewayController) fasthttp.RequestHandler {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		path := ctx.URI().PathOriginal()

		// GET
		if ctx.IsGet() {
			switch utils.ByteToString(path) {
			case "/":
				gc.Index(ctx)
			case "/apps":
				gc.Apps(ctx)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}

		// POST
		if ctx.IsPost() {
			switch utils.ByteToString(path) {
			case "/slack/shortcut":
				gc.SlackShortcut(ctx)
			case "/slack/command":
				gc.SlackCommand(ctx)
			case "/slack/event":
				gc.SlackEvent(ctx)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
