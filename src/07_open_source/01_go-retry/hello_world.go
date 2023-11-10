// @Time : 2023/2/17 17:27
// @Author : xiaoweiwei
// @File : hello_world

package main

import (
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

func main() {
	b := retry.NewExponential(1 * time.Second)

	// 最大重试次数
	// Stop after 4 retries, when the 5th attempt has failed. In this example, the worst case elapsed
	// time would be 1s + 1s + 2s + 3s = 7s.
	b = retry.WithMaxRetries(4, b)

	// 上线持续时间
	// Ensure the maximum value is 2s. In this example, the sleep values would be
	// 1s, 1s, 2s, 2s, 2s, 2s...
	//b = retry.WithCappedDuration(3 * time.Second, b)

	// 最大持续时间
	// Ensure the maximum total retry time is 5s.
	b = retry.WithMaxDuration(5*time.Second, b)

	for i := 0; i < 9; i++ {
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
