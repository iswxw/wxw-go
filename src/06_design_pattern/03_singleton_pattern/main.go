/*
@Time : 2022/2/24 11:25
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"sync"
)

type instance struct{}

var once sync.Once
var ins *instance

// 单例模式
func main() {
	fmt.Printf("%#v\n", getObj())
}

func getObj() *instance {
	if ins != nil {
		once.Do(func() {
			ins = &instance{}
		})
	}
	return ins
}
