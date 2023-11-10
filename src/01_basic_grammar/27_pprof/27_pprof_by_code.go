/*
@Time: 2022/12/22 11:03
@Author: wxw
@File: 27_pprof_by_code
*/
package main

import (
	"os"
	"runtime/pprof"
)

func main1() {
	cpuOut, _ := os.Create("cpu.out")
	defer cpuOut.Close()
	pprof.StartCPUProfile(cpuOut)
	defer pprof.StopCPUProfile()

	memOut, _ := os.Create("mem.out")
	defer memOut.Close()
	defer pprof.WriteHeapProfile(memOut)

	Sum(3, 5)

}

func Sum(a, b int) int {
	return a + b
}
