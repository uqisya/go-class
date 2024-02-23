package helpers

import "sync"

type Displayable interface {
	DisplayMsg(counter int, msg ...string)
	// DisplayMsgWg(wg *sync.WaitGroup, counter int, msg ...string)
	DisplayMsgMutex(wg *sync.WaitGroup, mtx *sync.Mutex, counter int, msg ...string)
}
