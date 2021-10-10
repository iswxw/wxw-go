/*
@Time: 2021/10/10 10:57
@Author: wxw
@File: demo
@Link: https://www.runoob.com/go/go-interfaces.html
*/
package main

import "fmt"

func main() {
	var phone Phone
	phone = new(XiaoMi)
	phone.call()

	phone = new(HuaWei)
	phone.call()
}

// Phone 定义一个手机接口
type Phone interface {
	call()
}

// XiaoMi 定义一个手机结构体
type XiaoMi struct{}

// Xiaomi方法
func (xiaomi XiaoMi) call() {
	fmt.Println("小米中该方法实现了Phone接口的方法")
}

// HuaWei 华为手机
type HuaWei struct{}

func (huawei HuaWei) call() {
	fmt.Println("华为中 实现了Phone 接口的call() 方法")
}
