package controllers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tsundata/assistant/internal/pkg/prometheus"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
)

func CreateInitControllersFn(gc *GatewayController) fasthttp.RequestHandler {
	fp := prometheus.New()
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		fp.Start()

		// GET
		if ctx.IsGet() {
			switch string(ctx.Path()) {
			case "/":
				gc.Index(ctx)
			case "/apps":
				gc.Apps(ctx)
			case "/metrics":
				fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
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

		fp.End(ctx)
	}

	return requestHandler
}
