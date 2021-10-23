/*
@Time: 2021/10/24 0:35
@Author: wxw
@File: demo02
*/
package main

import "fmt"

// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
func main() {
	var idata IData //接口变量
	var data1 data1 // 实例变量
	idata = data1   // 赋值
	idata.compute()
}

type IData interface {
	compute()
}

type data1 int

func (date data1) compute() {
	fmt.Println("非结构体也可以实现接口方法")
}
