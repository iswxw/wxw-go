/*
@Time: 2022/11/5 17:31
@Author: wxw
@File: demo01_hello
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(100)

	f := func(index int) {
		defer wg.Done()
		mu.Lock()

		fmt.Println(index)
		time.Sleep(time.Microsecond * 10)

		mu.Unlock()
	}

	for i := 100; i > 0; i-- {
		go f(i)
	}
	wg.Wait()
}
