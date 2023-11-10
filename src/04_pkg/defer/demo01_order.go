/*
@Time: 2022/2/20 17:02
@Author: wxw
@File: demo01
*/
package main

import "fmt"

// defer的执行顺序
// 遇见defer 依次入栈、函数返回依次出栈执行
func main() {
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
