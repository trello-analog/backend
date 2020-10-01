package emails

import (
	"fmt"
	"github.com/trello-analog/backend/services"
)

type EmailSender struct {
	EmailService *services.EmailService
}

func NewEmailSender() *EmailSender {
	service := services.NewEmailService()

	return &EmailSender{
		EmailService: service,
	}
}

func (es *EmailSender) SendEmailAfterSignUp(data *SignUpEmail) error {
	parseError := es.EmailService.ParseTemplate("emails/templates/SignUpEmail.html", data)
	subjectText := "Добро пожаловать в KindaTrello!"

	if parseError != nil {
		return parseError
	}

	sendEmail := es.EmailService.
		SetEmail(data.Email).
		SetSubject(subjectText).
		SendEmail()

	if sendEmail != nil {
		return sendEmail
	}
	fmt.Println("Letter has been sent!")
	return nil
}

func (es *EmailSender) SendForgotPasswordEmail(data *SignUpEmail) error {
	parseError := es.EmailService.ParseTemplate("emails/templates/ForgotPasswordEmail.html", data)
	subjectText := "Восстановление пароля"

	if parseError != nil {
		return parseError
	}

	sendEmail := es.EmailService.
		SetEmail(data.Email).
		SetSubject(subjectText).
		SendEmail()

	if sendEmail != nil {
		return sendEmail
	}

	return nil
}

func (es *EmailSender) SendTwoAuthCode(data *TwoAuthEmail) error {
	parseError := es.EmailService.ParseTemplate("emails/templates/TwoAuthCode.html", data)

	if parseError != nil {
		return parseError
	}

	sendEmail := es.EmailService.
		SetEmail(data.Email).
		SetSubject("Двухфакторная аутентификация").
		SendEmail()

	if sendEmail != nil {
		return sendEmail
	}

	return nil
}
