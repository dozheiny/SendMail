package main

import (
	"log"
	"net/smtp"
)

func Send(from string, to []string, password string, message []byte){

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	message = []byte("the body of email")

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Email Sent.")

}

func main() {

}
