/*
@Time: 2022/12/11 16:41
@Author: wxw
@File: demo04_value
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequest04(ctx)

	time.Sleep(10 * time.Second)
}

func HandelRequest04(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}
