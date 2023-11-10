/*
@Time : 2020/10/26 22:00
@Author : wxw
@File : 01_if_else
@Software: GoLand
*/
package main

import "fmt"

// if条件判断
func main() {
	age := 19

	if age > 18 { // true
		fmt.Println("君恩深似海")
	} else { // false
		fmt.Println("陈恩重如山")
	}

	// 多层if 嵌套
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("少年无为")
	} else {
		fmt.Println("好好学习，天天向上")
	}

	// 作用域的问题
	// age1 变量 此时只在if 条件判断语句中生效
	if age1 := 19; age1 > 18 {
		fmt.Println("君恩深似海")
	} else { // false
		fmt.Println("陈恩重如山")
	}
	// fmt.Println(age1) // 此处age1 已经失效 在作用域之外

}
