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
				memoRe := regexp.MustCompile(`^/memo/[\w\-]+$`)
				if memoRe.Match(path) {
					wc.Memo(ctx)
					return
				}
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
				credentialsRe := regexp.MustCompile(`^/credentials/[\w\-]+$`)
				if credentialsRe.Match(path) {
					wc.Credentials(ctx)
					return
				}
				credentialsCreateRe := regexp.MustCompile(`^/credentials/[\w\-]+/create$`)
				if credentialsCreateRe.Match(path) {
					wc.CredentialsCreate(ctx)
					return
				}
				settingRe := regexp.MustCompile(`^/setting/[\w\-]+$`)
				if settingRe.Match(path) {
					wc.Setting(ctx)
					return
				}
				settingCreateRe := regexp.MustCompile(`^/setting/[\w\-]+/create$`)
				if settingCreateRe.Match(path) {
					wc.SettingCreate(ctx)
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

		// POST
		if ctx.IsPost() {
			switch utils.ByteToString(path) {
			case "/":
				wc.Index(ctx)
			default:
				credentialsCreateRe := regexp.MustCompile(`^/credentials/[\w\-]+/store$`)
				if credentialsCreateRe.Match(path) {
					wc.CredentialsStore(ctx)
					return
				}
				settingCreateRe := regexp.MustCompile(`^/setting/[\w\-]+/store$`)
				if settingCreateRe.Match(path) {
					wc.SettingStore(ctx)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}
