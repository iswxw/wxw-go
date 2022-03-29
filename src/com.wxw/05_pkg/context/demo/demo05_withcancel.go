/*
@Time : 2022/3/29 16:31
@Author : weixiaowei
@File : demo05_withcancel
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// 控制单线程 通知 线程退出
	// funcOneThread()

	// 控制多线程
	funcTwoThread()
}

func funcOneThread() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask05(ctx, "worker1")

	time.Sleep(3 * time.Second)
	// 通知子协程退出，会执行 ctx.Done()
	cancel()

	// 等待主线程执行 下一句
	time.Sleep(3 * time.Second)
}

func funcTwoThread() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask05(ctx, "worker1")
	go reqTask05(ctx, "worker2")

	time.Sleep(3 * time.Second)
	// 通知子协程退出，会执行 ctx.Done()
	cancel()

	// 等待主线程执行 下一句
	time.Sleep(3 * time.Second)
}

func reqTask05(ctx context.Context, name string) {
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
