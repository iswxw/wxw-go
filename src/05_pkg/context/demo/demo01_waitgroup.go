package main

import (
	"fmt"
	"sync"

	"time"
)

// https://www.liwenzhou.com/posts/Go/go_context/
var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("All Task Done")
}
