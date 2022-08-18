package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

var (
	from      = "go-mail@noreply.com"
	user      = "2ac2981b8899b1"
	password  = "5e7784e948c122"
	host      = "smtp.mailtrap.io"
	port      = "2525"
	to        = []string{"john.doe@gmail.com", "joan.doe@gmail.com"}
	subject   = "The Test Mail"
	body      = "This is test mail by golang smtp."
	msgString = fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",from, strings.Join(to,","), subject, body)
	msg       = []byte(msgString)
)

func main() {
	auth := smtp.PlainAuth("", user, password, host)
	if err := smtp.SendMail(host+":"+port, auth, from, to, msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email Sent Successfully!")
	}
}
