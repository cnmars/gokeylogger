package main

import (
	"log"
	"os"

	"github.com/sevlyar/go-daemon"
)

var ctx *daemon.Context

func main() {
	err := checkRoot()
	checkErr(err)
	d, err := ctx.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()
	log.Println(os.Getpid())
	RunKeyLogger()
}

func init() {
	ctx = &daemon.Context{
		PidFileName: conf.GetString("pid_file"),
		PidFilePerm: 0755,
		LogFileName: conf.GetString("log_file"),
		LogFilePerm: 0755,
		WorkDir:     "./",
		Umask:       027,
	}
	logfile = OpenFile(conf.GetString("key_log_file"))
}
