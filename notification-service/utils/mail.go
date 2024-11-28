package utils

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(To, Subject, Content string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", viper.GetString("EMAIL"))
	mailer.SetHeader("To", To)
	mailer.SetHeader("Subject", Subject)
	mailer.SetBody("text/plain", Content)

	d := gomail.NewDialer("smtp.gmail.com", 587, viper.GetString("EMAIL"), viper.GetString("APP_PASSWORD"))

	if err := d.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
