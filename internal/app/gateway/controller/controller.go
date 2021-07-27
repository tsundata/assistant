package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"net/http"
	"strings"
	"time"
)

func CreateInitControllersFn(gc *GatewayController) func(router fiber.Router) {
	requestHandler := func(router fiber.Router) {
		defer func() {
			if err := recover(); err != nil {
				gc.logger.Error(err.(error))
			}
		}()

		// Middleware
		router.Use(recoverMiddleware.New())
		router.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))
		router.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "127.0.0.1"
			},
			Max:        500,
			Expiration: time.Minute,
		}))
		router.Use(newrelic.NewMiddleware(
			newrelic.MiddlewareConfig{
				NewRelicApp: gc.nr.Application(),
			},
		))

		router.Get("/", gc.Index)
		router.Post("/slack/event", gc.SlackEvent)
		router.Post("/telegram/event", gc.TelegramEvent)
		router.Post("/debug/event", gc.DebugEvent)

		// internal
		auth := func(c *fiber.Ctx) error {
			token := c.Get("Authorization")
			if token == "" {
				return c.SendStatus(http.StatusForbidden)
			}
			token = strings.ReplaceAll(token, "Bearer ", "")
			reply, err := gc.userSvc.Authorization(context.Background(), &pb.TextRequest{Text: token})
			if err != nil {
				gc.logger.Error(err)
				return c.SendStatus(http.StatusForbidden)
			}
			if !reply.GetState() {
				return c.SendStatus(http.StatusForbidden)
			}
			return c.Next()
		}
		router.Post("auth", gc.Authorization)
		router.Get("page", gc.GetPage)
		router.Post("webhook/trigger", gc.WebhookTrigger)
		router.Get("credential", gc.GetCredential)
		internal := router.Group("/").Use(auth)
		internal.Get("apps", gc.GetApps)
		internal.Post("app/oauth", gc.StoreAppOAuth)
		internal.Get("messages", gc.GetMessages)
		internal.Get("masking_credentials", gc.GetMaskingCredentials)
		internal.Post("credential", gc.CreateCredential)
		internal.Get("settings", gc.GetSettings)
		internal.Post("setting", gc.CreateSetting)
		internal.Get("action/messages", gc.GetActionMessages)
		internal.Post("action/message", gc.CreateActionMessage)
		internal.Delete("workflow/message", gc.DeleteWorkflowMessage)
		internal.Post("message/run", gc.RunMessage)
		internal.Post("message/send", gc.SendMessage)
		internal.Get("role/image", gc.GetRoleImage)

		router.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).SendString("Unsupported path")
		})
	}

	return requestHandler
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewGatewayController)
