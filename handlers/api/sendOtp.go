package api

import "github.com/gofiber/fiber/v2"

type PayloadSendOTP struct {
	Key string `json:"key" form:"key" validate:"required,min=5"`
	OTP string `json:"otp" form:"otp" validate:"required,len=4"`
}

func (a *Handler) SendOTP(c *fiber.Ctx) error {
	in, err := ReadJSON[PayloadSendOTP](c, c.Body())
	if err != nil {
		response := NewHTTPResponse(fiber.StatusBadRequest, err.Error(), "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	otp, err := a.Redis.Get(in.Key).Result()
	if err != nil {
		response := NewHTTPResponse(fiber.StatusBadRequest, "OTP code is expired, try to make new request instead", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if otp != in.OTP {
		response := NewHTTPResponse(fiber.StatusBadRequest, "incorrect OTP code", "")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var data = struct {
		Msg      string `json:"message"`
	}{
		Msg: "OTP has been sent successfully",
	}

	response := NewHTTPResponse(fiber.StatusOK, "", data)
	return c.Status(fiber.StatusOK).JSON(response)
}