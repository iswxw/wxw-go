/*
@Time: 2021/10/25 0:11
@Author: wxw
@File: demo01
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
