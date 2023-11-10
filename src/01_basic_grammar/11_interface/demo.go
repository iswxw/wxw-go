/*
@Time: 2021/10/10 10:57
@Author: wxw
@File: case
@Link: https://www.runoob.com/go/go-interfaces.html
*/
package main

import "fmt"

func main() {
	var phone IPhone
	phone = new(XiaoMi)
	phone.call()

	// 结构体变量，实现了call() 方法也就实现了IPhone接口
	var huawei HuaWei
	phone = huawei
	phone.call()
}

// IPhone 定义一个手机接口
type IPhone interface {
	call()
}

// XiaoMi 定义一个手机结构体
type XiaoMi struct{}

// Xiaomi方法   通过实现IPhone接口中call()方法 用XiaoMi实例指向它
func (xiaomi XiaoMi) call() {
	fmt.Println("小米中该方法实现了Phone接口的方法")
}

// HuaWei 华为手机
type HuaWei struct{}

func (huawei HuaWei) call() {
	fmt.Println("华为中 实现了Phone 接口的call() 方法")
}
