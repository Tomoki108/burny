package infrastructure

import (
	"fmt"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var client *ses.SES

func ConnectAWSSES() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Conf.AWS_REGION),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS SES session: %w", err)
	}
	client = ses.New(sess)
	return nil
}

const (
	sender  = "no-reply@burny.page"
	charSet = "UTF-8"
)

type AWSSESMailer struct {
}

func NewAWSSESMailer() domain.Mailer {
	return &AWSSESMailer{}
}

func (m *AWSSESMailer) Send(mail domain.Mail) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(mail.To)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(mail.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(mail.Subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err := client.SendEmail(input)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
