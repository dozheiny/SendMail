package main

import (
	"bufio"
	"fmt"
	"log"
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
	path := os.Args[2]
	password := os.Args[3]
	message := os.Args[4]

	// read .txt file to read and send emails
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	Send(from, fileTextLines, password, message)

	defer readFile.Close()

}
