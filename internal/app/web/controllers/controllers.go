package controllers

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
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

				qrRe := regexp.MustCompile(`^/qr/(.*)$`)
				if qrRe.Match(path) {
					wc.Qr(ctx)
					return
				}
				// auth
				if checkUUID(ctx.Path(), wc.midClient) {
					memoRe := regexp.MustCompile(`^/memo/[\w\-]+$`)
					if memoRe.Match(path) {
						wc.Memo(ctx)
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
				} else {
					ctx.Error("Forbidden", fasthttp.StatusForbidden)
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
				// auth
				if checkUUID(ctx.Path(), wc.midClient) {
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
				} else {
					ctx.Error("Forbidden", fasthttp.StatusForbidden)
					return
				}
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			}
		}
	}

	return requestHandler
}

func checkUUID(path []byte, midClient pb.MiddleClient) bool {
	uuid := extractUUID(path)
	if uuid == "" {
		return false
	}

	reply, err := midClient.Authorization(context.Background(), &pb.TextRequest{
		Text: uuid,
	})
	if err != nil {
		return false
	}

	return reply.State
}

func extractUUID(path []byte) string {
	re := regexp.MustCompile(`(\w{8}\-\w{4}\-\w{4}\-\w{4}\-\w{12})`)
	return re.FindString(utils.ByteToString(path))
}
