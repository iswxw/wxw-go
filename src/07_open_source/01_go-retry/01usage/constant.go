// @Time : 2023/2/17 17:23
// @Author : xiaoweiwei
// @File : main

package main

import (
	"fmt"
	"github.com/sethvargo/go-retry"
	"time"
)

// 常量增长
func main() {
	b := retry.NewConstant(1 * time.Second)

	for i := 0; i < 5; i++ {
		val, _ := b.Next()
		fmt.Printf("%v\n", val)
	}
	// Output:
	// 1s
	// 1s
	// 1s
	// 1s
	// 1s
}
