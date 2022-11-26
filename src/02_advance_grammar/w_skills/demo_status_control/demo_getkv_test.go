/*
@Time: 2021/11/28 15:52
@Author: wxw
@File: demo_getkv_test
*/
package main

import "testing"

// 基准测试
// 地址：https://mp.weixin.qq.com/s/MaovIWtQWVp9ZZL1DlppOQ
// D:\Project\wxw-go\src\com.wxw\02_advance_grammar\w_skills\demo_status_control>go test -bench=.
// goos: windows
// goarch: amd64
// 05_pkg: com.wxw/02_advance_grammar/w_skills/demo_status_control
// cpu: AMD Ryzen 7 4700U with Radeon Graphics
// BenchmarkSwitch-8       1000000000               0.2515 ns/op
// BenchmarkMap-8          100000000               11.21 ns/op
// PASS
// ok      com.wxw/02_advance_grammar/w_skills/demo_status_control     1.539s

// 结论：
// - switch 版本比 map 版本快了近 60 倍。此外，要较真的话，map 版本还用了一个 map 数据结构，占用额外的空间。
// - 性能差别这么大，其实通过汇编可以看到 map 版本调用了一个 runtime.mapaccess2 _ fast64(SB) 函数：而switch只是普通指令
func BenchmarkSwitch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OrderStateSwitch(0)
		OrderStateSwitch(1)
		OrderStateSwitch(2)
		OrderStateSwitch(3)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OrderStateMap(0)
		OrderStateMap(1)
		OrderStateMap(2)
		OrderStateMap(3)
	}
}
