package api

import (
	"otp-webapp/models/mailer"

	"github.com/rs/zerolog"
)

type Handler struct {
	Mailer mailer.Mailer
	Log    *zerolog.Logger
}