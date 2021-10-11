/*
@Time: 2021/10/12 0:04
@Author: wxw
@File: demo_recover_ok
*/
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("捕获到异常:", recover())
	}()
	panic("手动抛出异常")

	// 输出
	// 捕获到异常: 手动抛出异常
}
