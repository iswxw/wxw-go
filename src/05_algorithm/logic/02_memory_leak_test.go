// @Time : 2023/5/30 14:45
// @Author : xiaoweiwei
// @File : 02_memory_leak_test

package logic

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/* 内存/协程泄露举例
 *
 * 相关材料
 *  1. 什么情况下会出现内存泄露：https://www.yuque.com/iswxw/smdr49/gmkcgdubfzb82xwy#zk6QB
 *  2. goroutine 泄露场景及防止分析：https://blog.csdn.net/S_FMX/article/details/116172834
 **/

// TestMemoryLeak 内存/协程泄露举例
func TestMemoryLeak(t *testing.T) {

}

// 泄露场景一：(1) 非缓冲通道，缺少发送器，导致接收阻塞
func TestMissingSender(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 110
		// 因为 ch 一直没有接收数据，所以这个协程会阻塞在这里。
		fmt.Println("current value not receiver")
	}(ch)
	fmt.Println("=== RESULT execute finished")
}

// 泄露场景二：(2) 非缓冲通道，缺少接收器，导致发送阻塞
func TestMissingReceiver(t *testing.T) {
	ch := make(chan int)
	go func(ch chan int) {
		// 因为 ch 一直没有被发送数据，所以这个协程会阻塞在这里。
		val := <-ch
		fmt.Println("current receiver value is", val)
	}(ch)
	fmt.Println("=== RESULT execute finished")
}

// 泄露场景三：(3) 死锁。多个协程由于竞争资源导致死锁
func TestDeadLock(t *testing.T) {

	// 1. 请求超时阻塞（请求一直等待）
	// 2. 互斥锁忘记解锁
	funcNotUnlock()

	// 3. 同步锁使用不当(加锁和解锁数量不一致)等情况
	funcSyncUsedIllegal()

}

// 泄露场景四：(4) 创建协程的没有回收
func TestUnRecovery(t *testing.T) {

}

// ================ 内部方法 ================

func funcSyncUsedIllegal() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 3; i++ {
		wg.Done()
		fmt.Println("v = ", i)
	}
	wg.Wait()
}

func funcNotUnlock() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total: ", total)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			total += 1
		}()
	}
}
