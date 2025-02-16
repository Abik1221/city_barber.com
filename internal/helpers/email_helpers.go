package helpers

import (
	"fmt"
	"net/smtp"
)

// SendEmail sends an email using SMTP
func SendEmail(to, subject, body string) error {
	config := configs.LoadConfig()

	from := "your_email@example.com"
	password := config.EmailAPI
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}