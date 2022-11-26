/*
@Time : 2022/4/29 20:48
@Author : weixiaowei
@File : demo01_operator
*/
package main

import "fmt"

// 声明一个函数类型
type Printer func(contents string) (n int, err error)

// 普通函数
func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer

	// 声明的函数printToStd的签名与Printer的是一致的，因此前者是后者的一个实现，即使它们的名称以及有的结果名称是不同的。
	p = printToStd

	i, err := p("something")
	fmt.Println(i, err)
}
