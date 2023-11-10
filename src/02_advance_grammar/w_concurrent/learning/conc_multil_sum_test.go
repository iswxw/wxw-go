// @Time : 2023/3/16 15:58
// @Author : xiaoweiwei
// @File : conc

package learning

import (
	"fmt"
	"sync"
	"testing"
)

// TestMultiGoSum 多协程求和
func TestMultiGoSum(t *testing.T) {
	c := make(chan int, 150)
	total := 0
	wg := sync.WaitGroup{}
	wg.Add(10)

	//go compute(10, c, &wg)
	//
	for i := 0; i < 10; i++ {
		go compute(i*10, c, &wg)
	}
	wg.Wait()

	close(c)
	for i := range c {
		total += i
	}

	fmt.Println(total)

}

func compute(nums int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		//fmt.Println(nums + i)
		c <- 3
	}
}

// TestSplitGroup 将100分成10组
func TestSplitGroup(t *testing.T) {
	numMap := make(map[int][]int, 0)
	for i := 0; i < 100; i++ {
		key := i % 10
		curArray := numMap[key]
		if len(curArray) == 0 {
			curArray = []int{}
		}
		curArray = append(curArray, i)
		numMap[key] = curArray
	}
	fmt.Printf("%v", numMap)
}
