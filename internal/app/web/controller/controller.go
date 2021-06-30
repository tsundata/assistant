package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"net/http"
	"time"
)

func CreateInitControllersFn(wc *WebController) func(router fiber.Router) {
	requestHandler := func(router fiber.Router) {
		defer func() {
			if err := recover(); err != nil {
				wc.logger.Error(err.(error))
			}
		}()

		router.Get("/", wc.Index)
		router.Get("/echo", wc.Echo)
		router.Get("/Robots.txt", wc.Robots)
		router.Get("/page/:uuid", wc.Page)
		router.Get("/app/:category", wc.App)
		router.Get("/oauth/:category", wc.OAuth)
		router.Get("/qr/:text", wc.Qr)
		// webhook
		router.Get("/webhook/:flag", wc.Webhook)
		router.Post("/webhook/:flag", wc.Webhook)

		// auth middleware
		auth := func(c *fiber.Ctx) error {
			uuid := util.ExtractUUID(c.Path())
			if uuid == "" {
				return errors.New("error param")
			}

			// cache
			key := fmt.Sprintf("web:auth:%s", uuid)
			r, err := wc.rdb.Get(context.Background(), key).Result()
			var reply *pb.StateReply
			if err != nil {
				if errors.Is(err, redis.Nil) {
					reply, err = wc.gateway.Authorization(&pb.TextRequest{
						Text: uuid,
					})
					if err != nil {
						return err
					}
				} else {
					wc.logger.Error(err)
					return c.SendStatus(http.StatusForbidden)
				}
			}
			if r == "1" {
				wc.gateway.AuthToken(uuid)
				return c.Next()
			}

			if reply != nil && reply.GetState() {
				wc.gateway.AuthToken(uuid)
				wc.rdb.Set(context.Background(), key, "1", time.Hour)
				return c.Next()
			}

			wc.rdb.Set(context.Background(), key, "0", time.Hour)
			return c.SendStatus(http.StatusForbidden)
		}

		// auth Group
		authR := router.Group("/", auth).Use(auth)

		authR.Get("/memo/:uuid", wc.Memo)
		authR.Get("/apps/:uuid", wc.Apps)

		authR.Get("/credentials/:uuid", wc.Credentials)
		authR.Get("/credentials/:uuid/create", wc.CredentialsCreate)
		authR.Post("/credentials/:uuid/store", wc.CredentialsStore)

		authR.Get("/setting/:uuid", wc.Setting)
		authR.Get("/setting/:uuid/create", wc.SettingCreate)
		authR.Post("/setting/:uuid/store", wc.SettingStore)

		authR.Get("/action/:uuid", wc.Action)
		authR.Get("/action/:uuid/create", wc.ActionCreate)
		authR.Get("/action/:uuid/run", wc.ActionRun)
		authR.Post("/action/:uuid/store", wc.ActionStore)

		authR.Post("/workflow/:uuid/delete", wc.WorkflowDelete)

		authR.Post("/role/:uuid", wc.Role)
	}

	return requestHandler
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewWebController)
