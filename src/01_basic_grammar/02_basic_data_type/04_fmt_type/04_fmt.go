/*
@Time : 2020/10/26 0:00
@Author : wxw
@File : 04_fmt_type.04_fmt_type
@Software: GoLand
*/
package main

import "fmt"

// fmt 占位符
func main() {

	var n = 100
	// 查看类型
	fmt.Printf("%T \n", n)
	fmt.Printf("%v \n", n)
	fmt.Printf("%b \n", n)
	fmt.Printf("%d \n", n)
	fmt.Printf("%o \n", n)
	fmt.Printf("%x \n", n)
	fmt.Printf("%T \n", n)
	var s = "hello 年少无为"
	fmt.Printf("%s \n", s)
	fmt.Printf("%v \n", s)  // hello 年少无为
	fmt.Printf("%#v \n", s) // "hello 年少无为"
}
