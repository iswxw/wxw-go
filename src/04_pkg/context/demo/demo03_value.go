/*
@Time : 2022/3/29 14:03
@Author : weixiaowei
@File : demo03_value
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// 子协程中传递参数
type Options struct {
	Interval time.Duration
}

func reqTask03(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			fmt.Println(op.Interval)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1})

	go reqTask03(vCtx, "worker1")
	go reqTask03(vCtx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
