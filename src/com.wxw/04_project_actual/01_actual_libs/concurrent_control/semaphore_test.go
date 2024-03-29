package concurrent_control

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sema := New(3)
	var lock sync.Mutex
	arr := make([]int, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("concurrent_control.AvailablePermits : %d \n", sema.AvailablePermits())
		sema.Acquire() //数量不足，阻塞等待
		go func(i int) {
			defer sema.Release()
			lock.Lock()
			defer lock.Unlock()
			arr = append(arr, i)
			fmt.Println("concurrent_control")
			time.Sleep(time.Second)
		}(i)

	}

	_ = sema.Wait()
	fmt.Printf("concurrent_control.AvailablePermits : %d \n", sema.AvailablePermits())

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("---------- TestSemaphore done ----------")
}

func TestTrySemaphore(t *testing.T) {
	sema := New(3)

	for i := 0; i < 10; i++ {
		fmt.Printf("concurrent_control.AvailablePermits : %d \n", sema.AvailablePermits())
		if sema.TryAcquire() { //不阻塞等待
			go func() {
				defer sema.Release()
				fmt.Println("concurrent_control")
				time.Sleep(time.Second)
			}()
		}
	}

	_ = sema.Wait()
	fmt.Printf("concurrent_control.AvailablePermits : %d \n", sema.AvailablePermits())

	fmt.Println("---------- TestTrySemaphore done ----------")
}

func TestErrReturnSemaphore(t *testing.T) {
	sema := New(3)

	for i := 0; i < 3; i++ {
		sema.Add(func() error {
			// do some things
			fmt.Println("do some things")
			time.Sleep(time.Second)

			return nil
		})
	}

	sema.Add(func() error {
		// do some things
		fmt.Println("some error occur")
		time.Sleep(time.Millisecond * 200)
		return errors.New("occur error")
	})

	for i := 0; i < 3000; i++ {
		sema.Add(func() error {
			// do some things
			fmt.Println("do some things again")
			time.Sleep(time.Second)

			return nil
		})
	}

	err := sema.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------- TestErrRetrunSemaphore done ----------")
}
