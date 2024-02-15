package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailService interface {
	SendEmail(to, subject, content string, cc, bcc, attachFiles []string) error
}

type GmailService struct {
	name                string
	senderEmailAdrress  string
	senderEmailPassword string
}

func NewGmailService(name, senderEmailAdrress, senderEmailPassword string) EmailService {
	return &GmailService{
		name:                name,
		senderEmailAdrress:  senderEmailAdrress,
		senderEmailPassword: senderEmailPassword,
	}
}

func (sender *GmailService) SendEmail(to, subject, content string, cc, bcc, attachFiles []string) error {
	email := email.NewEmail()

	htmlContent, err := parseTemplate(to, "email", content)
	if err != nil {
		return err
	}

	email.From = fmt.Sprintf("%s <%s>", sender.name, sender.senderEmailAdrress)
	email.To = []string{to}
	email.Cc = cc
	email.Bcc = bcc
	email.Subject = subject
	email.HTML = []byte(htmlContent.Bytes())

	for _, f := range attachFiles {
		_, err := email.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s : %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.senderEmailAdrress, sender.senderEmailPassword, smtpAuthAddress)
	return email.Send(smtpServerAddress, smtpAuth)
}

func parseTemplate(email, templateType, content string) (*bytes.Buffer, error) {
	bodyTpl, err := template.ParseFiles(fmt.Sprintf("./static/%s.html", templateType))
	if err != nil {
		return nil, err
	}
	otp := rand.Intn(900000) + 100000
	var body bytes.Buffer
	data := map[string]string{"otp": strconv.Itoa(otp), "email": email, "content": content}
	if err := bodyTpl.Execute(&body, data); err != nil {
		return nil, err
	}
	return &body, nil
}
