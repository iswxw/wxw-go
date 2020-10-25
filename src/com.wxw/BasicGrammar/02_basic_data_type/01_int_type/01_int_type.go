/*
@Time : 2020/10/25 23:22
@Author : wxw
@File : 01_int_type
@Software: GoLand
*/
package main

import "fmt"

// 整型
func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d \n", i1)
	fmt.Printf("%b \n", i1) // 十进制数转换为2进制数
	fmt.Printf("%o \n", i1) // 十进制数转换为8进制数
	fmt.Printf("%x \n", i1) // 十进制数转换为16进制数
	// 八进制, 以 0开头
	i2 := 077
	fmt.Printf("%d \n", i2)
	// 十六进制，以 0x开头
	i3 := 0x1234567
	fmt.Printf("%d \n", i3)

	// 查看变量类型
	fmt.Printf("查看i3 变量类型: %T \n", i3)

	// 声明int8类型的变量（8位的整型变量）
	i4 := int8(8) // 明确指定int8,否则就默认为int型
	fmt.Printf("查看i4 变量类型: %T \n", i4)

}
