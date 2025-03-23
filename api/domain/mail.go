package domain

import "github.com/Tomoki108/burny/config"

type Mail struct {
	To      string
	Subject string
	Body    string
}

func NewEmailVerificationMail(to, token string) Mail {
	url := config.Conf.Host + "/api/v1/verify_email?token=" + token
	return Mail{
		To:      to,
		Subject: "Burny Email Verification",
		Body: "Please verify your email by clicking the link in 20 minitues: " + url +
			"\nIf you did not request this, please ignore this email.",
	}
}

type Mailer interface {
	Send(mail Mail) error
}
