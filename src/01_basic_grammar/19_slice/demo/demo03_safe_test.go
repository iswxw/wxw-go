/*
@Time: 2023/9/14 23:02
@Author: wxw
@File: demo03_safe_test
*/
package demo

import (
	"sync"
	"testing"
)

/**
* 切片非并发安全
* 多次执行，每次得到的结果都不一样
* 可以考虑使用 channel 本身的特性 (阻塞) 来实现安全的并发读写
 */
func TestSliceConcurrencySafe(t *testing.T) {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log("len = ", len(a))
	// not equal 10000
}

// 方式一：通过加锁实现slice线程安全，适合对性能要求不高的场景。
func TestSliceConcurrencySafeByMutex(t *testing.T) {
	var lock sync.Mutex //互斥锁
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			a = append(a, i)
		}(i)
	}
	wg.Wait()
	t.Log(len(a))
	// equal 10000
}

// 方式二：通过channel实现slice线程安全，适合对性能要求高的场景。
func TestSliceConcurrencySafeByChanel(t *testing.T) {
	buffer := make(chan int)
	a := make([]int, 0)

	// 消费者
	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()

	// 生产者
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer <- i
		}(i)
	}

	wg.Wait()
	t.Log(len(a))
	// equal 10000
}
