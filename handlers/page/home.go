package page

import "github.com/gofiber/fiber/v2"

func (p *Handler) Home(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("index", fiber.Map{
		"Title": "OTP WebApp",
	})
}