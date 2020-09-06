package services

import (
	"bytes"
	"github.com/trello-analog/backend/config"
	"html/template"
	"net/smtp"
)

type Email struct {
	subject string
	email   string
	body    string
}

type EmailService struct {
	auth  smtp.Auth
	email Email
}

func NewEmailService() *EmailService {
	cfg := config.GetConfig()
	auth := smtp.PlainAuth("", cfg.Email.Address, cfg.Email.Password, "smtp.gmail.com")
	return &EmailService{
		auth: auth,
	}
}

func (e *EmailService) ParseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	e.email.body = buf.String()
	return nil
}

func (e *EmailService) SetEmail(email string) *EmailService {
	e.email.email = email

	return e
}

func (e *EmailService) SetSubject(subject string) *EmailService {
	e.email.subject = "Subject: " + subject + "!\n"

	return e
}

func (e *EmailService) SendEmail() error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(e.email.subject + mime + "\n" + e.email.body)
	addr := "smtp.gmail.com:587"
	to := []string{e.email.email}

	err := smtp.SendMail(addr, e.auth, "info@kindatrello.com", to, msg)

	if err != nil {
		return err
	}

	return nil
}
