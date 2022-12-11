/*
@Time: 2022/12/11 16:38
@Author: wxw
@File: demo03_timeout
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandelRequest03(ctx)

	time.Sleep(10 * time.Second)
}

func HandelRequest03(ctx context.Context) {
	go WriteRedis03(ctx)
	go WriteDatabase03(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis03(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase03(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
