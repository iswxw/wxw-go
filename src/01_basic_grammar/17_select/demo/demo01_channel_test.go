/*
@Time : 2022/3/22 16:18
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	ch := make(chan int)

	c := 0
	stopCh := make(chan bool)

	go Chan(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive C", c)
		case s := <-ch:
			fmt.Println("Receive S", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

func Chan(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}
