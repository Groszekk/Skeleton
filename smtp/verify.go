package smtp

import (
	"BetterDash/models"
	"log"
	"net/smtp"
)

func Verify(user *models.UserData) {
	auth := smtp.PlainAuth("", "", "", "smtp.mailtrap.io")
	to := []string{user.Email}
	msg := []byte("From: \r\n" +
		"To: " + user.Email + "\r\n" +
		"Subject: Fucking Title\r\n" +
		"\r\n" +
		"Hi.\r\n" +
		"Welcome to the my platform. Click in the link down below to verify your e-mail address\r\n" +
		"http://localhost:8080/verify?token=" + user.Verify + "\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
