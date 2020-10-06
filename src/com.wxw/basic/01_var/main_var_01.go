/*
@Time : 2020/10/6 16:15
@Author : wxw
@File : main_01
@Software: GoLand
*/
package main

import "fmt"

// 批量声明
var (
	name string // ""
	age  int    //0
	isOk bool   //false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	// go语言中变量声明则必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) //%s：占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)              //打印完指定内容后，加一个换行符

	// 声明变量的同时赋值
	var s1 string = "声明变量的同时赋值"
	fmt.Println(s1)
	// 类型推到：根据值判断该变量的类型
	var s2 = "类型推到：根据值判断该变量的类型"
	println(s2)

	// 简短变量声明,只能在函数中使用
	s3 := "简短变量声明，只能在函数中使用"
	println(s3)

	// 匿名变量 特点是一个下画线“_”  匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
	a, _ := GetData()
	_, b := GetData()
	fmt.Printf("匿名变量 特点是一个下画线:%d,%d\n", a, b)
}

// 定义一个函数
func GetData() (int, int) {
	return 100, 200
}
