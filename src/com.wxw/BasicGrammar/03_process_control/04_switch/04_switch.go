/*
@Time : 2020/11/15 11:18
@Author : wxw
@File : 04_switch
@Software: GoLand
*/
package main

import "fmt"

// switch 简化大量判断（一个变量和某个具体的值做比较）
func main() {

	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效数字")
	}

	// switch 简化上面的代码
	switch x := 3; x {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效数字")
	}

	// switch 拓展 (并列过滤)
	switch y := 7; y {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
	// 条件判断
	age := 25
	switch {
	case age < 18:
		fmt.Println("少年宫")
	case age > 18:
		fmt.Println("青年才俊")
	default:
		fmt.Println("无效判断")
	}

}
