package controllers

import (
	"github.com/valyala/fasthttp"
	"log"
)

func CreateInitControllersFn(wc *WebController) fasthttp.RequestHandler {
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
				wc.Index(ctx)
			case "/Robots.txt":
				wc.Robots(ctx)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
