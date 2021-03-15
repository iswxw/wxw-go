package main

/**
 * @Description
 * @Author wxw
 * @link: https://www.liwenzhou.com/posts/Go/14_concurrence/
 * @Date 2021/3/14 18:57
 **/

import (
	"fmt"
	"sync"
)

// 定义变量
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine 就登记+1
		go hello1(i)
	}
	wg.Wait() // 等待所有的登记都结束
}

func hello1(i int) {
	defer wg.Done() // goroutine 结束就登记-1
	fmt.Println("hello goroutine", i)
}
