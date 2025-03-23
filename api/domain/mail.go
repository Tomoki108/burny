package domain

type Mail struct {
	To      string
	Subject string
	Body    string
}

func NewEmailVerificationMail(to, token string) Mail {
	return Mail{
		To:      to,
		Subject: "Burny Email Verification",
		Body: "Please verify your email by clicking the link in 20 minitues: https://example.com/verify?token=" + token +
			"\n\nIf you did not request this, please ignore this email.",
	}
}

type Mailer interface {
	Send(mail Mail) error
}
