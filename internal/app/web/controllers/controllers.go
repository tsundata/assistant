package controllers

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
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

		// auth
		authMiddleware := func(c *fiber.Ctx) error {
			uuid := utils.ExtractUUID(c.Path())
			if uuid == "" {
				return errors.New("error param")
			}

			reply, err := wc.midClient.Authorization(context.Background(), &pb.TextRequest{
				Text: uuid,
			})
			if err != nil {
				return err
			}

			if reply.GetState() {
				return c.Next()
			}
			return c.SendStatus(http.StatusForbidden)
		}
		router.Get("/memo/:uuid", authMiddleware, wc.Memo)
		router.Get("/apps/:uuid", authMiddleware, wc.Apps)

		router.Get("/credentials/:uuid", authMiddleware, wc.Credentials)
		router.Get("/credentials/:uuid/create", authMiddleware, wc.CredentialsCreate)
		router.Post("/credentials/:uuid/store", authMiddleware, wc.CredentialsStore)

		router.Get("/setting/:uuid", authMiddleware, wc.Setting)
		router.Get("/setting/:uuid/create", authMiddleware, wc.SettingCreate)
		router.Post("/setting/:uuid/store", authMiddleware, wc.SettingStore)

		router.Get("/scripts/:uuid", authMiddleware, wc.Scripts)
		router.Get("/scripts/:uuid/create", authMiddleware, wc.ScriptCreate)
		router.Post("/script/:uuid/store", authMiddleware, wc.ScriptStore)

		router.Get("/action/:uuid", authMiddleware, wc.Action)
		router.Get("/action/:uuid/create", authMiddleware, wc.ActionCreate)
		router.Get("/action/:uuid/run", authMiddleware, wc.ActionRun)
		router.Post("/action/:uuid/store", authMiddleware, wc.ActionStore)
	}

	return requestHandler
}
