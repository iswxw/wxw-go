/*
@Time : 2022/2/24 11:45
@Author : weixiaowei
@File : demo02_hungry
*/
package main

import "fmt"

// 饿汉模式
// 缺点：如果singleton创建初始化比较复杂耗时时，加载时间会延长。
func main() {
	hungry := GetInsHungry06()
	fmt.Println(hungry)
}

// Singleton 饿汉式单例
type singletonHungry06 struct{}

var insHungry06 *singletonHungry06

// 初始化
func init() {
	insHungry06 = &singletonHungry06{}
}

func GetInsHungry06() *singletonHungry06 {
	return insHungry06
}
