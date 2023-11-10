/*
@Time : 2022/2/24 11:39
@Author : weixiaowei
@File : demo01
*/
package main

import "fmt"

// 懒汉模式
// 非线程安全。当正在创建时，有线程来访问此时ins = nil就会再创建，单例类就会有多个实例了。
func main() {

	getIns := GetIns()
	fmt.Println(getIns)
}

type singleton struct{}

var ins *singleton

// 获取实例
func GetIns() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
