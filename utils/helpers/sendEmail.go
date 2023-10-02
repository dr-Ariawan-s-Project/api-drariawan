package helpers

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const BASE_URL_FRONTEND = "https://my-health-questionaire.netlify.app"

func SendMailQuestionerLink(email, codeAttempt, gmailPass string) {
	var subject_email = "Link Questioner dr. Ariawan App"
	var message = `
Terima kasih telah melakukan pendaftaran.
Berikut adalah link pengisian questioner anda:
`
	message = message + BASE_URL_FRONTEND + "/kuisioner?code=" + codeAttempt

	errSendEmail := sendMail(email, subject_email, message, gmailPass)
	if errSendEmail == nil {
		log.Println("success send email '" + subject_email + "' to " + email)
	} else {
		log.Println("error send email to "+email+". ", errSendEmail)
	}

}

func sendMail(email, subject, message, gmailPass string) error {
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "dr. Ariawan Questioner <fakhry.fun@gmail.com>"
	const CONFIG_AUTH_EMAIL = "fakhry.fun@gmail.com"
	var CONFIG_AUTH_PASSWORD = gmailPass

	if subject == "" || email == "" {
		return errors.New("subject and email must be fill")
	}

	var to = []string{email}
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
