package main

import (
	"github.com/robfig/cron"
)

var sendMailDone chan struct{}

func CronJob(d string) {
	sendMailDone = make(chan struct{}, 1)
	cr := cron.New()
	cr.AddFunc(d, SendMail)
	cr.Start()
	<-sendMailDone
}
