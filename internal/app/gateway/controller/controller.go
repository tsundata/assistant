package controller

import (
	"context"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/websocket/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	_ "github.com/tsundata/assistant/docs"
	"github.com/tsundata/assistant/internal/app/gateway/chat"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"log"
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
		router.Use(requestid.New(requestid.Config{
			ContextKey: enum.RequestIdKey,
		}))
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
		router.Use("/ws", func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			return fiber.ErrUpgradeRequired
		})

		// swagger
		router.Get("/swagger/*", swagger.Handler)

		// ws
		h := chat.NewHub(gc.bus, gc.logger, gc.chatbotSvc, gc.messageSvc)
		go h.Run()
		go h.EventHandle()
		router.Get("/ws/:uuid", websocket.New(func(conn *websocket.Conn) {
			room := conn.Params("uuid")
			log.Println(conn.Query("token"))
			chat.ServeWs(h, conn, room)
		}))

		// route
		router.Get("/", gc.Index)
		router.Post("/auth", gc.Authorization)
		router.Post("/webhook/trigger", gc.WebhookTrigger)
		router.Get("/health", gc.Health)

		// internal
		auth := func(c *fiber.Ctx) error {
			token := c.Get("Authorization")
			if token == "" {
				return c.SendStatus(http.StatusForbidden)
			}
			token = strings.ReplaceAll(token, "Bearer ", "")
			reply, err := gc.userSvc.Authorization(context.Background(), &pb.AuthRequest{Token: token})
			if err != nil {
				gc.logger.Error(err)
				return c.SendStatus(http.StatusForbidden)
			}
			if !reply.GetState() {
				return c.SendStatus(http.StatusForbidden)
			}
			c.Locals(enum.AuthKey, reply.Id)
			return c.Next()
		}
		internal := router.Group("/")
		internal.Use(auth)
		internal.Get("page", gc.GetPage)
		internal.Get("chart", gc.GetChart)
		internal.Get("apps", gc.GetApps)
		internal.Post("app/oauth", gc.StoreAppOAuth)
		internal.Get("messages", gc.GetMessages)
		internal.Get("masking_credentials", gc.GetMaskingCredentials)
		internal.Get("credential", gc.GetCredential)
		internal.Post("credential", gc.CreateCredential)
		internal.Get("settings", gc.GetSettings)
		internal.Post("setting", gc.CreateSetting)
		internal.Get("action/messages", gc.GetActionMessages)
		internal.Post("action/message", gc.CreateActionMessage)
		internal.Delete("workflow/message", gc.DeleteWorkflowMessage)
		internal.Post("message/run", gc.RunMessage)
		internal.Post("message/send", gc.SendMessage)
		internal.Get("role/image", gc.GetRoleImage)

		// 404
		router.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).SendString("Unsupported path")
		})
	}

	return requestHandler
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewGatewayController)
