/*
@Time: 2022/2/20 17:25
@Author: wxw
@File: demo06_panic_recover
*/
package main

import "fmt"

// defer遇见panic，并捕获异常, 后面的defer也会被执行
func main() {
	deferCall01()

	fmt.Println("main 正常结束")
}

func deferCall01() {

	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

	panic("异常内容") //触发defer出栈

	defer func() {
		fmt.Println("defer: panic 之后, 永远执行不到")
	}()
}
