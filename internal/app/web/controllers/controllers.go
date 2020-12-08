package controllers

import (
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
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
				pageRe := regexp.MustCompile(`^/page/[\w\-]+$`)
				if pageRe.Match(ctx.Path()) {
					wc.Page(ctx)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
