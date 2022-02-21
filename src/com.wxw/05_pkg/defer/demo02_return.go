/*
@Time: 2022/2/20 17:05
@Author: wxw
@File: demo02_return
*/
package main

import "fmt"

// defer与return谁先谁后
// return之后的语句先执行，defer后的语句后执行
func main() {
	returnAndDefer()
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}
