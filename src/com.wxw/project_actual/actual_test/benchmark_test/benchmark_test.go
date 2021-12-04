/*
@Time: 2021/12/4 23:43
@Author: wxw
@File: benchmark_test
*/
package main

import (
	"fmt"
	"testing"
)

// 基准测试 测试代码 case_1
func BenchmarkHandleType(b *testing.B) {
	flag := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println("BenchMark Test:", HandleType(flag))
	}
}
