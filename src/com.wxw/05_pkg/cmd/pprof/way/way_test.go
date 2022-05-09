/*
@Time : 2022/5/9 14:34
@Author : weixiaowei
@File : way_test
*/
package way

import "testing"

/**
  - 生成文件 执行指令：go test -bench=".*" -cpuprofile cpu.profile -memprofile mem.profile
  - CPU 性能分析调用关系文件生成： go tool pprof -png cpu.profile > cpu.png
  - 浏览器查看文件：go tool pprof -http="0.0.0.0:8081" way.test cpu.profile
*/
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(i, 1)
	}
}
