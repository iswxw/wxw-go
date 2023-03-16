// @Time : 2023/2/1 15:46
// @Author : xiaoweiwei
// @File : main_test

package conc

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"sync"
	"sync/atomic"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var count atomic.Int64
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			count.Add(1)
			fmt.Println(count.Load())
		})
	}
	wg.Wait()
}

func TestNums(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
