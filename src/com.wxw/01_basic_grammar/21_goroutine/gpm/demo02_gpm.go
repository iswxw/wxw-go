/*
@Time : 2022/4/30 11:26
@Author : weixiaowei
@File : demo02_gpm
*/
package main

import (
	"fmt"
	"sync/atomic"
)

// 打印结果分析：https://time.geekbang.org/column/article/39841
// 结论是：不会有任何内容被打印出来
func main() {
	// notOrder()

	order()
}

// 顺序执行 goroutine
func order() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}

func notOrder() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	// time.Sleep(time.Millisecond * 500)
}
