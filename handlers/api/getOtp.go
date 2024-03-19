package api

import (
	"otp-webapp/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (a *Handler) GetOTP(c *fiber.Ctx) error {
	key := utils.GenerateRandomString(30)
	otp := utils.GenerateOTP(4)
	expire := 2 * time.Minute

	err := a.Redis.Set(key, otp, expire).Err()
	if err != nil {
		response := NewHTTPResponse(fiber.StatusInternalServerError, err.Error(), "")
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	err = a.Mailer.SendOTP("habi@ternaklinux.com", otp)
	if err != nil {
		a.Log.Error().Str("error", err.Error()).Msg("failed to send otp to email")
	}

	var data = struct {
		Msg      string `json:"message"`
		Key			string `json:"key"`
	}{
		Msg: "OTP sent to your email",
		Key: key,
	}

	response := NewHTTPResponse(fiber.StatusOK, "", data)
	return c.Status(fiber.StatusOK).JSON(response)
}