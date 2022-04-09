/*
@Time : 2022/2/26 12:07
@Author : weixiaowei
@File : demo02_once
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	getIns05 := GetIns05()
	fmt.Println(getIns05)
}

// GetLazyInstance 懒汉式
type singleton05 struct{}

var (
	ins05  *singleton05
	once05 sync.Once
)

func GetIns05() *singleton05 {
	if ins05 == nil {
		once05.Do(func() {
			ins05 = &singleton05{}
		})
	}
	return ins05
}
