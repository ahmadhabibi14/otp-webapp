package mailer

import (
	"github.com/rs/zerolog"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func NewMailer(l *zerolog.Logger) Mailer {
	var mlr Mailer = Mailer{}
	mh, err := NewMailhog()
	if err != nil {
		l.Error().Str("error", err.Error()).Msg(`cannot load mailhog`)
	} else {
		mlr.SendMailFunc = mh.SendEmail
	}

	return mlr
}

func (m *Mailer) SendOTP(email, otp string) error {
	return m.SendMailFunc(
		map[string]string{email: ``},
		`OTP Code`,
		`OTP Code is `+otp,
		`<p>OTP Code is `+otp+`</p>`,
	)
}
