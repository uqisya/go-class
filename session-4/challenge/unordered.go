package main

import (
	"session-4/challenge/helpers"
	"session-4/challenge/model"
	"time"
)

func main() {

	var data1 helpers.Displayable = &model.Data{}

	for i := 1; i <= 4; i++ {
		go data1.DisplayMsg(i, "coba1", "coba2", "coba3")
		go data1.DisplayMsg(i, "bisa1", "bisa2", "bisa3")
	}
	time.Sleep(time.Second * 2)

}
