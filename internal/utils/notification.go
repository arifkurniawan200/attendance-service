package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypeSms   NotificationType = "sms"
)

type Notification struct {
	Type    NotificationType
	Subject string
	Body    string
	Target  string
}

func SendNotification(notification Notification) error {
	if notification.Type == NotificationTypeEmail {
		// Set up the email sender
		email := "arifkurniawandev96@gmail.com"
		password := "opsoacmvrhazvgyn"

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", email)
		mailer.SetHeader("To", notification.Target)
		mailer.SetHeader("Subject", notification.Subject)
		mailer.SetBody("text/plain", notification.Body)

		// Create a new SMTP client session
		dialer := gomail.NewDialer("smtp.gmail.com", 587, email, password)

		// Send the email
		if err := dialer.DialAndSend(mailer); err != nil {
			return err
		}
		fmt.Println("Email sent successfully!")

	} else if notification.Type == NotificationTypeSms {
		// do action if have env for send sms
	}
	return nil
}
