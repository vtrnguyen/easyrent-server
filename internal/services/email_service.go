package services

import (
	"bytes"

	"easyrent-server/configs"
	email "easyrent-server/internal/shared/job"

	"text/template"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	config configs.MailConfig
}

// NewEmailService creates a new instance of EmailService with the necessary dependencies.
func NewEmailService() *EmailService {
	return &EmailService{
		config: configs.GetMailConfig(),
	}
}

// Send sends an email based on the provided EmailJob. It uses a template to generate the email body and sends the email using the configured SMTP server.
func (s *EmailService) Send(
	job email.Job,
) error {
	tmpl, err := template.ParseFiles(
		"internal/shared/templates/" + job.Template,
	)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, job.Data); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader(
		"From",
		s.config.From,
	)
	m.SetHeader(
		"To",
		job.To,
	)
	m.SetHeader(
		"Subject",
		job.Subject,
	)
	m.SetBody(
		"text/html",
		body.String(),
	)

	d := gomail.NewDialer(
		s.config.Host,
		587,
		s.config.Username,
		s.config.Password,
	)

	err = d.DialAndSend(m)

	if err != nil {
		return err
	}

	return nil
}