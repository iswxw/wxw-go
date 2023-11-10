/*
@Time : 2020/10/6 16:39
@Author : wxw
@File : main_const_02
@Software: GoLand
*/
package main

import "fmt"

// 声明常量
const Pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一直
const (
	n1 = 100
	n2
	n3
)

// iota：go语言常量计数器，只能在常量表达式中使用
/**
 * iota在const关键字出现时将被重置为0。
 * const中【每新增一行常量声明】将使iota计数一次(加1)
 * 在定义枚举时很有用
 */
const (
	x1 = iota //0
	x2        //1
	x3        //2
	x4        //3
)

const (
	a1 = iota //0
	a2 = 100  //100
	a3 = iota //2
	a4        //3
)
const a5 = iota //0

// << 表示左移 做网盘
const (
	_  = iota             // # 这里用的匿名变量，默认值 0开始
	KB = 1 << (10 * iota) //这里 iota 数值是1 10*1 =10
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//Pi = 123
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)
	fmt.Println("x3:", x3)
	fmt.Println("x4:", x4)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)
	fmt.Println("a5:", a5)
}
