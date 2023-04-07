// @Time : 2023/4/7 14:47
// @Author : xiaoweiwei
// @File : demo02_async_preempt

package main

import "fmt"

// https://zhuanlan.zhihu.com/p/386998235
// https://zhuanlan.zhihu.com/p/387003228
func main() {
	go func(n int) {
		for {
			n++
			fmt.Println(n)
		}
	}(0)
	for {
	}
}
