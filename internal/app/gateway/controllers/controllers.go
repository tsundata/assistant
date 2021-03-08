package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func CreateInitControllersFn(gc *GatewayController) func(router fiber.Router) {

	requestHandler := func(router fiber.Router) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		router.Get("/", gc.Index)
		router.Post("/slack/event", gc.SlackEvent)

		router.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).SendString("Unsupported path")
		})
	}

	return requestHandler
}
