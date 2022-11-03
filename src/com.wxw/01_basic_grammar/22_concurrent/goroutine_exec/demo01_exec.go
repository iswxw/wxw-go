/*
@Time : 2022/4/30 11:26
@Author : weixiaowei
@File : demo02_gpm
*/
package main

import (
	"fmt"
)

// 打印结果分析：https://time.geekbang.org/column/article/39841
// 结论是：不会有任何内容被打印出来
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	// time.Sleep(time.Millisecond * 500)
}
