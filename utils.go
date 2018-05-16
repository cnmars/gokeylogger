package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var logfile *os.File
var killfile *os.File

func WriteKillFile() {
	killfile.WriteString(fmt.Sprintln("#!/bin/bash"))
	killfile.WriteString(fmt.Sprintf("sudo kill %d \n", os.Getpid()))
	killfile.WriteString(fmt.Sprintln("sudo rm pid"))
	killfile.Sync()
	killfile.Close()
}
func ReadLogFile(f string) string {
	b, err := ioutil.ReadFile(f)
	checkErr(err)
	return string(b)
}
func OpenFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	checkErr(err)
	return file
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func init() {
	logfile = OpenFile(conf.GetString("key_log_file"))
	killfile = OpenFile("kill.sh")
}
