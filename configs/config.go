package configs

import "os"

type MailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func GetMailConfig() MailConfig {
	return MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		From:     os.Getenv("MAIL_FROM"),
	}
}