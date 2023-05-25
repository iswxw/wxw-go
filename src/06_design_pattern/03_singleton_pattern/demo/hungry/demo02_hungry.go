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
	hungry := GetInsHungry()
	fmt.Println(hungry)
}

type singletonHungry struct{}

var insHungry *singletonHungry = &singletonHungry{}

func GetInsHungry() *singletonHungry {
	return insHungry
}
