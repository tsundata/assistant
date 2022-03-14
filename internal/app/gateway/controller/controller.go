package controller

import (
	"context"
	"errors"
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
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
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
		router.Use(recoverMiddleware.New(recoverMiddleware.Config{EnableStackTrace: true}))
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

		// swagger
		router.Get("/swagger/*", swagger.New())

		// WebSocket
		router.Use("/ws", func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			return fiber.ErrUpgradeRequired
		})
		h := chat.NewHub(gc.bus, gc.logger, gc.messageSvc)
		go h.Run()
		go h.EventHandle()
		router.Get("/ws/group/:uuid", websocket.New(func(conn *websocket.Conn) {
			ctx := context.Background()
			// auth
			token := conn.Query("token")
			authReply, err := gc.userSvc.Authorization(ctx, &pb.AuthRequest{Token: token})
			if err != nil {
				gc.logger.Error(err)
				_ = conn.WriteMessage(websocket.TextMessage, []byte("Forbidden"))
				return
			}
			if !authReply.GetState() {
				_ = conn.WriteMessage(websocket.TextMessage, []byte("Forbidden"))
				return
			}

			// group id
			uuid := conn.Params("uuid")
			ctx = md.BuildAuthContext(authReply.Id)
			groupReply, err := gc.chatbotSvc.GetGroup(ctx, &pb.GroupRequest{Group: &pb.Group{Uuid: uuid}})
			if err != nil {
				_ = conn.WriteMessage(websocket.TextMessage, []byte("Error group"))
				return
			}

			chat.ServeWs(h, conn, groupReply.Group.Id, authReply.Id)
		}))
		router.Get("/ws/notify", websocket.New(func(conn *websocket.Conn) {
			// auth
			token := conn.Query("token")
			userId, err := getUser(gc.userSvc, token)
			if err != nil {
				_ = conn.WriteMessage(websocket.TextMessage, []byte("Forbidden"))
				return
			}

			gc.Notify(conn, userId)
		}))

		// route
		router.Get("/", gc.Index)
		router.Get("/Robots.txt", gc.Robots)
		router.Post("/auth", gc.Authorization)
		router.Post("/webhook/trigger", gc.WebhookTrigger)
		router.Get("/health", gc.Health)
		router.Get("/page", gc.GetPage)
		router.Get("/app/:category", gc.App)
		router.Get("/oauth/:category", gc.OAuth)
		router.Get("/qr/:text", gc.QR)
		router.Get("/webhook/:flag", gc.Webhook)
		router.Post("/webhook/:flag", gc.Webhook)

		// internal
		auth := func(c *fiber.Ctx) error {
			// auth
			token := c.Get("Authorization")
			userId, err := getUser(gc.userSvc, token)
			if err != nil {
				gc.logger.Error(err)
				return c.SendStatus(http.StatusForbidden)
			}
			c.Locals(enum.AuthKey, userId)
			return c.Next()
		}
		internal := router.Group("/")
		internal.Use(auth)

		internal.Get("groups", gc.GetGroups)
		internal.Get("group", gc.GetGroup)
		internal.Get("messages", gc.GetMessages)
		internal.Get("message", gc.GetMessage)
		internal.Get("user", gc.GetUser)
		internal.Get("group/setting", gc.GetGroupSetting)
		internal.Post("group/setting", gc.UpdateGroupSetting)
		internal.Get("group/bots", gc.GetGroupBots)
		internal.Get("group/bot", gc.GetGroupBot)
		internal.Get("group/bot/setting", gc.GetGroupBotSetting)
		internal.Post("group/bot/setting", gc.UpdateGroupBotSetting)
		internal.Get("inboxes", gc.GetInboxes)
		internal.Get("user/setting", gc.GetUserSetting)
		internal.Get("system/setting", gc.GetSystemSetting)
		internal.Post("setting", gc.UpdateSetting)
		internal.Get("apps", gc.GetApps)

		// 404
		router.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).SendString("Unsupported path")
		})
	}

	return requestHandler
}

func getUser(userSvc pb.UserSvcClient, token string) (int64, error) {
	token = strings.ReplaceAll(token, "Bearer ", "")
	reply, err := userSvc.Authorization(context.Background(), &pb.AuthRequest{Token: token})
	if err != nil {
		return 0, errors.New("forbidden")
	}
	if !reply.GetState() {
		return 0, errors.New("forbidden")
	}

	return reply.Id, nil
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewGatewayController)
