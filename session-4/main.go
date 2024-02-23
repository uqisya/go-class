package main

import (
	"fmt"
	"time"
)

func main() {
	go printName("routine 1")
	go printName("routine 2")
	go printName("routine 3")
	defer printName("routine 4")
	go printName("routine 5")
	time.Sleep(time.Second * 5)
}

func printName(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		// time.Sleep(time.Second * 5)
	}
}
