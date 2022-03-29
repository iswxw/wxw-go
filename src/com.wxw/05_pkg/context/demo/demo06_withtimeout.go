/*
@Time : 2022/3/29 16:42
@Author : weixiaowei
@File : demo06_withtimeout
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// 控制子协程执行时间
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTask06(ctx, "worker1")
	go reqTask06(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}

func reqTask06(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}
