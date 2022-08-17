/*
@Time: 2021/9/21 17:45
@Author: wxw
@File: case
*/
package main

import "fmt"

func main() {
	var a int = 19 // 申明一个实际变量
	var b *int     // 声明一个指针变量

	if b == nil {
		fmt.Println("当前指针为空")
	}

	b = &a // 指针变量存储的地址

	fmt.Printf("值：%d    变量的地址：%x    类型：%T \n", a, &a, a)

	// 使用指针访问变量
	fmt.Printf("值：%d    变量的地址：%x    类型：%T \n", *b, b, b)

	// 空指针
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

}
