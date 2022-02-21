/*
@Time: 2022/2/20 17:09
@Author: wxw
@File: demo03_function_init
*/
package main

import "fmt"

// 函数的返回值初始化
func main() {
	testInitReturnValue(3)
}

func testInitReturnValue(v int) (t int) {
	fmt.Println("t = ", t)
	return 2
}
