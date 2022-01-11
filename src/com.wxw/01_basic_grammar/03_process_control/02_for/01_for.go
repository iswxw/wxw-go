/*
@Time : 2020/10/26 22:11
@Author : wxw
@File : 01_for
@Software: GoLand
*/
package main

import "fmt"

// for 循环
func main() {

	// 基本表达式
	for i := 0; i < 10; i++ {
		fmt.Print(i, "、")
	}

	// 变种1——初始值可以省略
	var i = 5
	for ; i < 10; i++ {
		fmt.Print(i, "、")
	}
	// 变种2——结束语句可以省略
	var j = 5
	for j < 10 {
		fmt.Print(j, "、")
		j++
	}

	// 无限循环
	//for {
	//	fmt.Print("123","、")
	//}

	// 键值循环 for...range
	// 数组、切片字符串返回索引和值
	// map 返回键和值
	// 通道（channel） 只返回通道内的内容
	s := "helllo 年少无为"
	for i, v := range s {
		fmt.Printf("(%d,%c)\n", i, v)
	}
}
