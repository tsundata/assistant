package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
	"time"
)

func CreateInitControllersFn(wc *WebController) func(router fiber.Router) {
	requestHandler := func(router fiber.Router) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		router.Get("/", wc.Index)
		router.Get("/echo", wc.Echo)
		router.Get("/Robots.txt", wc.Robots)
		router.Get("/page/:uuid", wc.Page)
		router.Get("/app/:category", wc.App)
		router.Get("/oauth/:category", wc.OAuth)
		router.Get("/qr/:text", wc.Qr)

		// auth middleware
		auth := func(c *fiber.Ctx) error {
			uuid := utils.ExtractUUID(c.Path())
			if uuid == "" {
				return errors.New("error param")
			}

			// cache
			key := fmt.Sprintf("web:auth:%s", uuid)
			s := wc.rdb.Get(context.Background(), key)
			r, err := s.Result()
			var reply *pb.StateReply
			if err != nil && err == redis.Nil {
				reply, err = wc.midClient.Authorization(context.Background(), &pb.TextRequest{
					Text: uuid,
				})
				if err != nil {
					wc.logger.Error(err)
					wc.rdb.Set(context.Background(), key, "0", time.Hour)
					return c.SendStatus(http.StatusForbidden)
				}
			}
			if r == "1" {
				return c.Next()
			}

			if reply.GetState() {
				wc.rdb.Set(context.Background(), key, "1", time.Hour)
				return c.Next()
			}

			wc.rdb.Set(context.Background(), key, "0", time.Hour)
			return c.SendStatus(http.StatusForbidden)
		}

		router.Get("/memo/:uuid", auth, wc.Memo)
		router.Get("/apps/:uuid", auth, wc.Apps)

		router.Get("/credentials/:uuid", auth, wc.Credentials)
		router.Get("/credentials/:uuid/create", auth, wc.CredentialsCreate)
		router.Post("/credentials/:uuid/store", auth, wc.CredentialsStore)

		router.Get("/setting/:uuid", auth, wc.Setting)
		router.Get("/setting/:uuid/create", auth, wc.SettingCreate)
		router.Post("/setting/:uuid/store", auth, wc.SettingStore)

		router.Get("/scripts/:uuid", auth, wc.Scripts)
		router.Get("/scripts/:uuid/create", auth, wc.ScriptCreate)
		router.Get("/script/:uuid/run", auth, wc.ScriptRun)
		router.Post("/script/:uuid/store", auth, wc.ScriptStore)

		router.Get("/action/:uuid", auth, wc.Action)
		router.Get("/action/:uuid/create", auth, wc.ActionCreate)
		router.Get("/action/:uuid/run", auth, wc.ActionRun)
		router.Post("/action/:uuid/store", auth, wc.ActionStore)

		router.Post("/workflow/:uuid/delete", auth, wc.WorkflowDelete)

		// webhook
		router.Get("/webhook/:flag", wc.Webhook)
		router.Post("/webhook/:flag", wc.Webhook)
	}

	return requestHandler
}
