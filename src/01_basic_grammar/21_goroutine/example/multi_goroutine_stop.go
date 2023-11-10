// @Time : 2023/3/16 19:46
// @Author : xiaoweiwei
// @File : multi_goroutine_stop

package main

import (
	"fmt"
	"sync"
)

// go v1.19.3
func main() {
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func(i int) {
			defer wg.Done()
			if i == 10 {
				panic(any("it will stop all goroutine"))
			}
			fmt.Println("i = ", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finished ")
}

//i =  0
//i =  8
//i =  4
//i =  2
//i =  3
//i =  6
//i =  18
//i =  5
//panic: it will stop all goroutine
