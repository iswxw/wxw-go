// @Time : 2023/2/1 15:46
// @Author : xiaoweiwei
// @File : main_test

package conc

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"sync/atomic"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var count atomic.Int64
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			count.Add(1)
		})
	}
	wg.Wait()

	fmt.Println(count.Load())
}
