/*
@Time : 2020/10/25 23:40
@Author : wxw
@File : 02_float_type
@Software: GoLand
*/
package main

import "fmt"

// 浮点数
func main() {

	// math.MaxFloat32() // float32 最大值
	f1 := 1.23456
	fmt.Printf("f1 的数据类型: %T \n", f1) // 默认go语言中小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("f2 显式申明float32类型: %T \n", f2) //显式申明float32类型
	// f1 = f2 // float32类型的值不能直接赋值给float64类型的变量
}
