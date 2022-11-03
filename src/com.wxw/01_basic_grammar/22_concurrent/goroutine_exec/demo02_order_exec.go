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
// 顺序执行 goroutine
func main() {
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
			fn := func() { fmt.Println(i) }
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}
