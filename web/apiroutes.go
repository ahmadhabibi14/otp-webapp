package web

import (
	"otp-webapp/handlers/api"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App, apis *api.Handler) {
	api := app.Group("/api")

	_ = api
}
