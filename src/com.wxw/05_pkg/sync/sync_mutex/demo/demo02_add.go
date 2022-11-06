/*
@Time: 2022/11/5 23:18
@Author: wxw
@File: demo02_race
*/
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	sync.Mutex
}

func (c *counter) add() {
	c.Lock()
	defer c.Unlock()

	c.count++
}

func (c *counter) value() int {
	c.Lock()
	defer c.Unlock()

	return c.count
}

func main() {
	var wg sync.WaitGroup
	var c counter

	wg.Add(2)

	// goroutine 1
	go func() {
		defer wg.Done()

		for i := 0; i < 5000; i++ {
			c.add()
		}
	}()

	// goroutine 2
	go func() {
		defer wg.Done()

		for i := 0; i < 5000; i++ {
			c.add()
		}
	}()

	wg.Wait()

	fmt.Println(c.value())
}
