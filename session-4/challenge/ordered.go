package main

import (
	"session-4/challenge/helpers"
	"session-4/challenge/model"
	"sync"
)

func main() {

	var data1 helpers.Displayable = &model.Data{}

	var wg sync.WaitGroup
	var mtx sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go data1.DisplayMsgMutex(&wg, &mtx, i, "coba1", "coba2", "coba3")
		go data1.DisplayMsgMutex(&wg, &mtx, i, "bisa1", "bisa2", "bisa3")
		wg.Wait()
	}
	// time.Sleep(time.Second * 2)

}

/*
	masih ada kebingungan penerapan hanya sekedar waitGroup aja kadang bisa udah ke-ordered (terurut)

	ada contoh dari modul mas novalagung:
	https://dasarpemrogramangolang.novalagung.com/A-mutex.html

	perlu dilakukan penerapan mutex, jika hanya waitGroup tetap muncul race condition

*/
