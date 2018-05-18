package main

import (
	"github.com/robfig/cron"
)

func CronJob(d string) {
	cr := cron.New()
	cr.AddFunc(d, SendMail)
	cr.Start()
}
