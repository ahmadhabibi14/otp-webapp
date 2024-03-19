package api

import (
	"errors"
	"net/http"
	"otp-webapp/models/mailer"
	"otp-webapp/utils"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Handler struct {
	Mailer mailer.Mailer
	Log    *zerolog.Logger
	Redis *redis.Client
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewHTTPResponse(code int, errors string, data any) HTTPResponse {
	return HTTPResponse{
		Code:   code,
		Status: http.StatusText(code),
		Errors: errors,
		Data:   data,
	}
}

func ReadJSON[T any](c *fiber.Ctx, b []byte) (T, error) {
	var body T
	err := c.BodyParser(&body)
	if err != nil {
		return body, errors.New("invalid Payload")
	}

	errvalid := utils.ValidateStruct(body)
	if errvalid != nil {
		return body, errvalid
	}

	return body, nil
}