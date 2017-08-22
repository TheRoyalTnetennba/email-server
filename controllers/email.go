package main

import (
	"log"
	"net/http"
	"net/smtp"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("Message")
	from := FromEmail
	pass := EmailPass
	to := ToEmail

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + r.Host + "\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("Email was sent")
}
