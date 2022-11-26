/*
@Time: 2022/2/20 17:24
@Author: wxw
@File: demo05_panic_not_process
*/
package main

import "fmt"

// defer遇见panic，但是并不捕获异常的情况
func main() {
	deferCall()

	fmt.Println("main 正常结束")
}

func deferCall() {
	defer func() { fmt.Println("defer: panic 之前1") }()
	defer func() { fmt.Println("defer: panic 之前2") }()

	panic("异常内容") //触发defer出栈

	defer func() {
		fmt.Println("defer: panic 之后，永远执行不到")
	}()
}
