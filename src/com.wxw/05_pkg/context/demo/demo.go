package main

import (
	"fmt"
	"sync"

	"time"
)

// https://www.liwenzhou.com/posts/Go/go_context/
var wg sync.WaitGroup

// 初始的例子

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("over")
}
