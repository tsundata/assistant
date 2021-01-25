package controllers

import (
	"github.com/tsundata/assistant/internal/pkg/utils"
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

		path := ctx.URI().PathOriginal()

		// GET
		if ctx.IsGet() {
			switch utils.ByteToString(path) {
			case "/":
				wc.Index(ctx)
			case "/Robots.txt":
				wc.Robots(ctx)
			default:
				pageRe := regexp.MustCompile(`^/page/[\w\-]+$`)
				if pageRe.Match(path) {
					wc.Page(ctx)
					return
				}
				appsRe := regexp.MustCompile(`^/apps/[\w\-]+$`)
				if appsRe.Match(path) {
					wc.Apps(ctx)
					return
				}
				qrRe := regexp.MustCompile(`^/qr/(.*)$`)
				if qrRe.Match(path) {
					wc.Qr(ctx)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
