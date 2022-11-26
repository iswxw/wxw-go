/*
@Time: 2022/11/5 17:30
@Author: wxw
@File: demo01_hello
*/
package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type Mutex struct {
	state int32 // 锁状态 0未锁/1已锁
}

func (m *Mutex) Lock() {
	for {
		// 原子cas操作
		if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
			break
		}
		// 睡眠一秒
		time.Sleep(time.Second)
	}
}

func (m *Mutex) Unlock() {
	// 重复解锁或者未锁状态解锁报异常
	if !atomic.CompareAndSwapInt32(&m.state, 1, 0) {
		panic("lock state error")
	}
}

func main() {
	var wg sync.WaitGroup
	var mu Mutex
	wg.Add(100)

	f := func(index int) {
		defer wg.Done()
		mu.Lock()
		time.Sleep(time.Microsecond * 10)
		mu.Unlock()
	}

	for i := 100; i > 0; i-- {
		go f(i)
	}

	wg.Wait()
}
