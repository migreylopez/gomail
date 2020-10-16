package gomail

import (
	"errors"
	"fmt"
	"net/smtp"
)

var emailAuth smtp.Auth

type EmailData struct {
	From     string
	Host     string
	Password string
	Port     string
}

func SendEmailSMTP(to []string, subject string, contentType string, data interface{}, templatePath string, emailData EmailData) (bool, error) {
	emailHost := emailData.Host
	emailFrom := emailData.From
	emailPassword := emailData.Password
	emailPort := emailData.Port

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(templatePath, data)
	if err != nil {
		return false, errors.New("unable to parse email template: " + err.Error())
	}

	mime := "MIME-version: 1.0;\nContent-Type: " + contentType + "; charset=\"UTF-8\";\n\n"
	subjectLine := "Subject: " + subject + "\n"
	msg := []byte(subjectLine + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}
