/*
@Time: 2022/3/12 22:47
@Author: wxw
@File: dbg_test
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count : ", count)
	}
}

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}
