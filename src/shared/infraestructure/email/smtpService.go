package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	mailservice "github.com/alejogs4/hn-website/src/shared/infraestructure/email/mailService"
)

// SMPTService concrete implementation of mailservice.Service to send emails using built in smtp golang package
type SMPTService struct{}

// Send function to finally send email to user or users
func (s SMPTService) Send(mail mailservice.Mail) error {
	// SMTP information
	smtpHost := "smtp.sendgrid.net" // Change this for a valid one
	smtpPort := "587"

	// SMTP Auth
	username := os.Getenv("SENDGRID_USER")
	password := os.Getenv("SENDGRID_PASSWORD")
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Parse passed template
	temp, err := template.ParseFiles(mail.Template)
	if err != nil {
		return err
	}

	// Add headers and subject
	mimeHeaders := make(map[string]string)
	mimeHeaders["From"] = mail.From
	mimeHeaders["Subject"] = mail.Subject
	mimeHeaders["MIME-Version"] = "1.0"
	mimeHeaders["Content-Type"] = "text/html; charset=\"utf-8\""

	headersMessage := ""
	for k, v := range mimeHeaders {
		headersMessage += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	var emailBody bytes.Buffer
	emailBody.Write([]byte(fmt.Sprintf("%s\n\n", headersMessage)))

	err = temp.Execute(&emailBody, mail.Body)
	if err != nil {
		return err
	}

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, mail.From, mail.To, emailBody.Bytes())
}
