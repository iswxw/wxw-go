/*
@Time : 2022/3/29 16:58
@Author : weixiaowei
@File : demo02_pprof_file
@link: https://geektutu.com/post/hpg-pprof.html
*/
package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	// 启动 prof 性能分析程序
	if err := pprof.StartCPUProfile(f); err != nil {
		return
	}
	defer pprof.StopCPUProfile()

	// 开始执行代码程序
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
