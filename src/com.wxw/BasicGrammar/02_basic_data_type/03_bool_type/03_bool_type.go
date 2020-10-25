/*
@Time : 2020/10/25 23:57
@Author : wxw
@File : 03_bool_type
@Software: GoLand
*/
package main

import "fmt"

// bool值
func main() {

	b1 := true
	var b2 bool // bool值默认是false,而且不能强转
	fmt.Printf("%T \n", b1)
	fmt.Printf("%T value:%v \n", b2)
}
