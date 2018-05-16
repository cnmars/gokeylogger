package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

var logfile *os.File

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
func checkRoot() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	if u.Uid != "0" {
		return fmt.Errorf("Cannot read device files. Are you running as root?")
	}
	return nil
}
func init() {
	logfile = OpenFile(conf.GetString("key_log_file"))
}
