/*
@Time : 2022/3/29 11:50
@Author : weixiaowei
@File : demo01_slice_array
@数据和切片的陷阱问题：https://geektutu.com/post/hpg-gotchas-array-slice.html
*/
package main

import "fmt"

func main() {
	// 问题一：类型传递和值传递 修改是否生效
	// funcValueAndTypeParams()

	// 问题二：扩容导致 元素修改失败
	funcCapacity()
}

func funcCapacity() {
	// 问题二：切片扩容导致修改值失效问题
	a01 := []int{1, 2}
	foo001(a01)
	fmt.Println("a = ", a01)

	// 方案一：设置返回值，将新切片返回并赋值给 main 函数中的变量 a。
	a02 := foo002(a01)
	fmt.Println("a = ", a02)

	// 方案二：切片也使用指针方式传参
	foo003(&a01)
	fmt.Println("a = ", a01)
}

func foo001(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
}

func foo002(a []int) []int {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
	return a
}

func foo003(a *[]int) {
	*a = append(*a, 1, 2, 3, 4, 5, 6, 7, 8)
	(*a)[0] = 200
}

func funcValueAndTypeParams() {
	// 问题一：数组
	a1 := [2]int{1, 2}
	foo1(a1) // 输出 [1 2]
	fmt.Println("a=", a1)

	// 解法一：切片
	a2 := []int{1, 2}
	foo2(a2) // 输出 [200 2]
	fmt.Println("a=", a2)

	// 解法二：指针+数组
	a3 := [2]int{1, 2}
	foo3(&a3) // 输出 [200 2]
	fmt.Println("a=", a3)
}

// 入参是数组
func foo1(a [2]int) {
	a[0] = 200
}

// 入参是切片
func foo2(a []int) {
	a[0] = 200
}

// 入参是 数组类型的指针
func foo3(a *[2]int) {
	(*a)[0] = 200
}
