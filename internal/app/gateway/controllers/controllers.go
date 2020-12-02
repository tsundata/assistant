package controllers

import (
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

		// GET
		if ctx.IsGet() {
			switch string(ctx.Path()) {
			case "/":
				gc.Index(ctx)
			case "/foo":
				gc.Foo(ctx)
			case "/apps":
				gc.Apps(ctx)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}

		// POST
		if ctx.IsPost() {
			switch string(ctx.Path()) {
			case "/slack/shortcut":
				gc.SlackShortcut(ctx)
			case "/slack/command":
				gc.SlackCommand(ctx)
			case "/slack/event":
				gc.SlackEvent(ctx)
			case "/slack/webhook":
				gc.AgentWebhook(ctx)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
