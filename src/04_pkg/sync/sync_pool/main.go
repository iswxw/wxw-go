/*
@Time : 2022/3/22 17:10
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{}
	fmt.Println(pool)
}
