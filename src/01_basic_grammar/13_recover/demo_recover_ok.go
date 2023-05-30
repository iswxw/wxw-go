/*
@Time: 2021/10/12 0:04
@Author: wxw
@File: demo_recover_ok
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	deferDemo1()
	recoverDemo1()

}

// 异常捕获 测试方法二
func recoverDemo1() {
	defer func() {
		if r := recover(); r != any(nil) {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
	panic(any("test panic"))

}

// 捕捉异常测试方法一
func deferDemo1() {

	defer func() {
		fmt.Println("捕获到异常:", recover())
	}()
	panic(any("手动抛出异常"))

	// 输出
	// 捕获到异常: 手动抛出异常
}
