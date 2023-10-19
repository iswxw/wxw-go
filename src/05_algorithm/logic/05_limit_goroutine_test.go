// @Time : 2023/10/17 20:09
// @Author : xiaoweiwei
// @File : 05_limit_goroutine_test

package logic

import (
	"fmt"
	"testing"
	"time"
)

// limitChan 限制协程个数
func Test05(t *testing.T) {
	limitChan := make(chan struct{}, 10) // 最大协程数限制为10个
	for i := 0; i < 100; i++ {
		go job(i, limitChan)
	}

}

// 任务函数
func job(i int, limitChan chan struct{}) {
	limitChan <- struct{}{}
	defer func() {
		<-limitChan
	}()

	// 执行任务
	time.Sleep(1 * time.Second)

	// 统计
	fmt.Printf("任务:%d已完成，当前协程数:%d\n", i, len(limitChan))
}

/**
参考材料
  1.限制协成个数：https://blog.csdn.net/qq_37102984/article/details/121791579

*/
