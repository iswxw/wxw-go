/*
@Time: 2022/2/20 17:18
@Author: wxw
@File: demo04_return_defer
*/
package main

import "fmt"

// 先return 后 defer
func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

	defer func() {
		t = t * 10
	}()

	return 1
}

func main() {
	fmt.Println(returnButDefer())
}
