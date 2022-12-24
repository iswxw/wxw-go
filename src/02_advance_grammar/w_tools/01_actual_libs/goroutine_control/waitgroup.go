// @Time : 2022/8/5 15:51
// @Author : xiaoweiwei
// @File : waitgroup

package sugar

import "sync"

type WaitGroup struct {
	wait sync.WaitGroup
	ch   chan bool
}

func NewWaitGroup(maxGoroutine int) *WaitGroup {
	var wait sync.WaitGroup
	ch := make(chan bool, maxGoroutine)
	return &WaitGroup{ch: ch, wait: wait}
}

func (w *WaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		w.ch <- true
		w.wait.Add(1)
	}
}

func (w *WaitGroup) Done() {
	w.wait.Done()
	<-w.ch
}

func (w *WaitGroup) Wait() {
	defer close(w.ch)
	w.wait.Wait()
}
