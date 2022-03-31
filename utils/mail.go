package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(name, email, subject, fileName, token string) {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	template := ParseHtml(fileName, map[string]string{
		"to":    email,
		"token": token,
	})

	err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, "from@example.com", []string{email}, []byte(template))
	if err != nil {
		fmt.Println(err)
	}
}
