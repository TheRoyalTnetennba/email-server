package main

import (
	"log"
	"net/http"
	"net/smtp"

	simplejson "github.com/bitly/go-simplejson"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("Message")
	from := FromEmail
	pass := EmailPass
	to := ToEmail

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + r.Header.Get("Origin") + "\n\n" +
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

func SmartBiz(w http.ResponseWriter, r *http.Request) {
	req, _ := simplejson.NewFromReader(r.Body)
	from := FromEmail
	pass := EmailPass
	to := ToEmail
	mail := req.Get("mail").MustString()
	message := req.Get("message").MustString()
	name := req.Get("name").MustString()
	body := "From: " + name + "\n" + "Message: " + message + "\n" + "Mail: " + mail

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"CC: " + TempCC + "\n" +
		"Subject: New Form Submission" + "\n\n" +
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
