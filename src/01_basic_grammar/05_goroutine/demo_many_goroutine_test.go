package _goroutine

/**
 * @Description
 * @Author wxw
 * @link: https://www.liwenzhou.com/posts/Go/14_concurrence/
 * @Date 2021/3/14 18:57
 **/

import (
	"fmt"
	"sync"
	"testing"
)

// 定义变量
var wgS sync.WaitGroup

func TestMany(t *testing.T) {

	for i := 0; i < 10; i++ {
		wgS.Add(1) // 启动一个goroutine 就登记+1
		go hello1(i)
	}
	wgS.Wait() // 等待所有的登记都结束
}

func hello1(i int) {
	defer wgS.Done() // goroutine 结束就登记-1
	fmt.Println("hello goroutine", i)
}
