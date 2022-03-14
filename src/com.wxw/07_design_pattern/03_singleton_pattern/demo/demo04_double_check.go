/*
@Time : 2022/2/26 11:54
@Author : weixiaowei
@File : demo04_double_check
*/
package main

import (
	"fmt"
	"sync"
)

// 缺点：避免了每次加锁，提高代码效率
func main() {
	getIns04 := GetIns04()
	fmt.Println(getIns04)
}

type singleton04 struct{}

var ins04 *singleton04
var mu04 sync.Mutex

func GetIns04() *singleton04 {
	if ins04 == nil {
		mu04.Lock()
		defer mu04.Unlock()
		if ins04 == nil {
			ins04 = &singleton04{}
		}
	}
	return ins04
}
