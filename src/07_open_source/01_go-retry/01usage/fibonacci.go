// @Time : 2023/2/17 17:29
// @Author : xiaoweiwei
// @File : fibonacci

package main

import (
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

// 斐波那契式增长
func main() {
	b := retry.NewFibonacci(1 * time.Second)

	for i := 0; i < 5; i++ {
		val, _ := b.Next()
		fmt.Printf("%v\n", val)
	}
	// Output:
	// 1s
	// 2s
	// 3s
	// 5s
	// 8s
}
