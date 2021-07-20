package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"net/http"
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

			reply, err := wc.gateway.Authorization(&pb.TextRequest{
				Text: uuid,
			})
			if err != nil {
				return err
			}

			if reply != nil && reply.GetState() {
				wc.gateway.AuthToken(uuid)
				return c.Next()
			}

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

		authR.Get("/role/:uuid", wc.Role)
	}

	return requestHandler
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewWebController)
