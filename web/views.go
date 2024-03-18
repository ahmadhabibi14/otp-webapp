package web

import (
	"otp-webapp/handlers/page"

	"github.com/gofiber/fiber/v2"
)

func WebViews(app *fiber.App, page *page.Handler) {
	app.Get("/", page.Home)
}