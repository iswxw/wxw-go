// @Time : 2023/4/6 14:31
// @Author : xiaoweiwei
// @File : demo02_race_goroutine

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
	"time"
)

// 异步抢占的例子：https://zhuanlan.zhihu.com/p/216118842
func main() {
	runtime.GOMAXPROCS(1)
	for {
		time.Sleep(time.Second * 5)
		go func() {
			fmt.Printf("tid %d\n", unix.Getpid())
			for {
			}
		}()
	}
	time.Sleep(time.Minute * 3)
}
