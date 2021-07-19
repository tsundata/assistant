package http

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestNewServer(t *testing.T) {
	_, err := New(nil, func(r fiber.Router) {}, nil)
	if err != nil {
		t.Fatal(err)
	}
}
