package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// smtp host configuration.
const smtpHost = "smtp.gmail.com"
const smtpPort = "587"

func Send(from string, to []string, password string, message string) {

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending Email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		log.Fatal("Error sending mail:", err)
	} else {
		fmt.Println("Email Sent !")
	}
}

func main() {

	// get command line arguments
	var arguments = make(map[string]string)
	arguments["from"] = os.Args[1]
	arguments["path"] = os.Args[2]
	arguments["password"] = os.Args[3]
	arguments["message"] = os.Args[4]

	// read .txt file to read and send emails
	readFile, err := os.Open(arguments["path"])
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	Send(arguments["from"], fileTextLines, arguments["password"], arguments["message"])

}
