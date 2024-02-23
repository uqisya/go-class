package model

import (
	"fmt"
	"sync"
)

type Data struct {
	text [3]string
}

func (d Data) DisplayMsg(counter int, msg ...string) {
	d.text[0] = msg[0]
	d.text[1] = msg[1]
	d.text[2] = msg[2]
	fmt.Printf("[%s %s %s] %d\n", d.text[0], d.text[1], d.text[2], counter)
}

// func (d Data) DisplayMsgWg(wg *sync.WaitGroup, counter int, msg ...string) {
// 	d.text[0] = msg[0]
// 	d.text[1] = msg[1]
// 	d.text[2] = msg[2]
// 	fmt.Printf("[%s %s %s] %d\n", d.text[0], d.text[1], d.text[2], counter)
// 	wg.Done()
// }

func (d *Data) DisplayMsgMutex(wg *sync.WaitGroup, mtx *sync.Mutex, counter int, msg ...string) {
	mtx.Lock()
	d.text[0] = msg[0]
	d.text[1] = msg[1]
	d.text[2] = msg[2]
	fmt.Printf("[%s %s %s] %d\n", d.text[0], d.text[1], d.text[2], counter)
	mtx.Unlock()
	wg.Done()
}
