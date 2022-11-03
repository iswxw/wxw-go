/*
@Time: 2022/11/4 0:13
@Author: wxw
@File: demo02_chan
*/
package main

import "fmt"

// https://github.com/hyper0x/Golang_Puzzlers/blob/master/src/puzzlers/article16/q2/demo39.go
func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
