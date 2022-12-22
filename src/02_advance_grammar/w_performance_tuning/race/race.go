/*
@Time: 2022/12/22 9:43
@Author: wxw
@File: race
*/
package main

import "time"

func main() {
	var x int
	go func() {
		for {
			x = 1
		}
	}()

	go func() {
		for {
			x = 2
		}
	}()
	time.Sleep(10 * time.Second)
}
