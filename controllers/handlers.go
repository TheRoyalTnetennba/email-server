package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/smtp"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	body := string(requestDump)
	from := FromEmail
	pass := EmailPass
	to := "gpaye8@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
