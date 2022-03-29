/*
@Time : 2022/3/29 16:03
@Author : weixiaowei
@File : demo02_map
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	o := sync.Once{}
	o.Do(func() {
		fmt.Println("")
	})
}
