package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "go-mail@noreply.com"
	user := "2ac2981b8899b1"
	password := "5e7784e948c122"

	// Receiver email address.
	to := []string{
		"t0239184ps@gmail.com",
	}

	// smtp server configuration.
	host := "smtp.mailtrap.io"
	port := "2525"

	// Message.
	msg := []byte("From: john.doe@example.com\r\n" +
		"To: roger.roe@example.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", user, password, host)

	// Sending email.
	err := smtp.SendMail(host+":"+port, auth, from, to, msg)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
