/*
@Time : 2022/3/29 14:36
@Author : weixiaowei
@File : demo04_deadline
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
// Deadline() (deadline time.Time, ok bool)

func reqTask04(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 5秒后取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go reqTask04(ctx, "worker1")
	go reqTask04(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
