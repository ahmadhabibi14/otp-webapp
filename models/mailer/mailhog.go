package mailer

import (
	"fmt"
	"otp-webapp/configs"

	"github.com/wneessen/go-mail"
)

type Mailhog struct {
	configs.MailhogConf
	client *mail.Client
}

func NewMailhog() (*Mailhog, error) {
	res := &Mailhog{
		MailhogConf: configs.EnvMailhog(),
	}
	err := res.Connect()
	return res, err
}

func (m *Mailhog) Connect() error {
	zlog := configs.InitLogger()
	if m.client != nil {
		err := m.client.Close()
		if err != nil {
			zlog.Error().Str("ERROR", err.Error()).
				Msg("Connection close")
		}
	}
	var err error
	m.client, err = mail.NewClient(m.MailhogHost,
		mail.WithPort(m.MailhogPort),
		mail.WithTLSPolicy(mail.NoTLS),
	)
	return err

}

var ErrMailhogSendingEmail = fmt.Errorf(`Mailhog SendEmail`)

func (m *Mailhog) SendEmail(toEmailName map[string]string, subject, text, html string) error {
	msg := mail.NewMsg()
	if err := msg.FromFormat(m.DefaultFromName, m.DefaultFromEmail); err != nil {
		return fmt.Errorf("%w: FromFormat: %v", ErrMailhogSendingEmail, err)
	}
	if m.UseBcc {
		if err := msg.AddToFormat(m.DefaultFromName, m.DefaultFromEmail); err != nil {
			return fmt.Errorf("%w: AddToFormat: %v", ErrMailhogSendingEmail, err)
		}
	}
	if err := msg.ReplyToFormat(m.DefaultFromName, m.ReplyToEmail); err != nil {
		return fmt.Errorf("%w: ReplyToFormat: %v", ErrMailhogSendingEmail, err)
	}
	for email, name := range toEmailName {
		if m.UseBcc {
			if err := msg.AddBccFormat(name, email); err != nil {
				return fmt.Errorf("%w: AddBccFormat: %v", ErrMailhogSendingEmail, err)
			}
		} else {
			if err := msg.AddToFormat(name, email); err != nil {
				return fmt.Errorf("%w: AddToFormat: %v", ErrMailhogSendingEmail, err)
			}
		}
	}
	msg.Subject(subject)
	if text != `` {
		msg.SetBodyString(mail.TypeTextPlain, text)
	}
	if html != `` {
		msg.SetBodyString(mail.TypeTextHTML, html)
	}
	if err := m.client.DialAndSend(msg); err != nil {
		return fmt.Errorf("%w: DialAndSend: %v", ErrMailhogSendingEmail, err)
	}
	return nil
}
