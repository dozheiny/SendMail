package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func Send(from string, to []string, password string, message string) {
	// smtp host configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// Sending Email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println("Error sending mail:", err)
	} else {
		fmt.Println("Email Sent !")
	}

}

func main() {

	// get command line arguments
	from := os.Args[1]
	to := os.Args[2]
	password := os.Args[3]
	message := os.Args[4]
	// send email
	Send(from, to, password, message)
}
