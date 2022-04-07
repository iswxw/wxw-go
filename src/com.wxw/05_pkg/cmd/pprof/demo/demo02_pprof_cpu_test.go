/*
@Time : 2022/3/29 16:58
@Author : weixiaowei
@File : demo02_pprof_file
@link: https://geektutu.com/post/hpg-pprof.html
*/
package main

import "testing"

const url = "hello world"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error")
	}
}

// go test -bench=. -cpuprofile=cpu.prof
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
