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
		`OTP Code Authentication`,
		`Your OTP code is `+otp+`

		Please use this OTP within the next 2 minutes to complete your authentication process. After this time, the OTP will expire, and you will need to request a new one.
		`,
		`<p>Your OTP code is <b>`+otp+`</b></p>
		<p>Please use this OTP within the next 2 minutes to complete your authentication process. After this time, the OTP will expire, and you will need to request a new one.</p>`,
	)
}
