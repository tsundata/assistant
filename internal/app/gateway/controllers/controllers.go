package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	"log"
	"net/http"
	"time"
)

func CreateInitControllersFn(gc *GatewayController) func(router fiber.Router) {
	requestHandler := func(router fiber.Router) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
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
			Max:        20,
			Expiration: time.Minute,
		}))

		router.Get("/", gc.Index)
		router.Post("/slack/event", gc.SlackEvent)
		router.Post("/telegram/event", gc.TelegramEvent)

		// internal
		internal := router.Group("/")
		internal.Get("page", gc.GetPage)

		router.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).SendString("Unsupported path")
		})
	}

	return requestHandler
}

var ProviderSet = wire.NewSet(CreateInitControllersFn, NewGatewayController)
