// @Time : 2023/2/17 17:28
// @Author : xiaoweiwei
// @File : exponential

package main

import (
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

// 指数增长
func main() {
	b := retry.NewExponential(1 * time.Second)
	b = retry.WithMaxRetries(3, b) // 最大重试次数

	for i := 0; i < 5; i++ {
		val, _ := b.Next()
		fmt.Printf("%v\n", val)
	}
	// Output:
	// 1s
	// 2s
	// 4s
	// 8s
	// 16s
}
