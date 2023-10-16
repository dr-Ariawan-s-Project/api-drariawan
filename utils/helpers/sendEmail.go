package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
)

// Request struct
type EmailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *EmailRequest {
	return &EmailRequest{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func SendMailQuestionerLink(email, codeAttempt, gmailPass string) {
	var baseUrlFrontend = config.InitConfig().BASE_URL_FE
	// Working Directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var subjectEmail = "Link Questioner dr. Ariawan App"
	r := NewRequest([]string{email}, subjectEmail, "-")

	templateData := struct {
		UrlLink string
	}{
		UrlLink: baseUrlFrontend + "/kuisioner?code=" + codeAttempt,
	}
	// 	var message = `
	// Terima kasih telah melakukan pendaftaran.
	// Berikut adalah link pengisian questioner anda:
	// `
	// 	message = message + baseUrlFrontend + "/kuisioner?code=" + codeAttempt
	errTemplate := r.ParseEmailTemplate(wd+"/utils/files/index.html", templateData)
	if errTemplate == nil {
		errSendEmail := r.SendMail(gmailPass)
		if errSendEmail == nil {
			log.Println("success send email '" + subjectEmail + "' to " + email)
		} else {
			log.Println("error send email to "+email+". ", errSendEmail)
		}
	} else {
		log.Println("error load template", errTemplate)
	}

}

func (r *EmailRequest) ParseEmailTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *EmailRequest) SendMail(gmailPass string) error {
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "dr. Ariawan App <drariawan.app@gmail.com>"
	const CONFIG_AUTH_EMAIL = "drariawan.app@gmail.com"
	var CONFIG_AUTH_PASSWORD = gmailPass

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	if r.subject == "" {
		return errors.New("subject must be fill")
	}

	var to = r.to
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + r.subject + "\n\n" +
		mime + "\n" + r.body

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, r.to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
