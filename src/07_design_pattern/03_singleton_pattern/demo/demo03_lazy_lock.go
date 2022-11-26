/*
@Time : 2022/2/26 11:38
@Author : weixiaowei
@File : demo03_lazy_lock
*/
package main

import (
	"fmt"
	"sync"
)

// 缺点：虽然解决并发的问题，但每次加锁是要付出代价的
func main() {
	lazyLock := GetInsLazyLock()
	fmt.Println(lazyLock)
}

type singletonLazyLock struct{}

var insLazyLock *singletonLazyLock
var mu sync.Mutex

func GetInsLazyLock() *singletonLazyLock {
	mu.Lock()
	defer mu.Unlock()

	if insLazyLock == nil {
		insLazyLock = &singletonLazyLock{}
	}
	return insLazyLock
}
