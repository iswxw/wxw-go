/*
@Time: 2021/10/24 23:40
@Author: wxw
@File: demo_cpu
*/
package _goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCPU(t *testing.T) {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)

	// 可以设置使用多个CPU
	gomaxprocs := runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("gomaxprocs = ", gomaxprocs)
}
