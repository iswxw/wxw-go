// @Time : 2023/10/17 19:05
// @Author : xiaoweiwei
// @File : 04_order_print_string_test

package logic

import (
	"fmt"
	"sync"
	"testing"
)

// 启三个协成 按顺 cat、fish、dog 分别打印100次

func Test04(t *testing.T) {
	wg := &sync.WaitGroup{}
	flagCat, flagFish, flagDog := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)

	wg.Add(3)

	flagCat <- struct{}{}
	go printCat(wg, flagCat, flagFish)
	go printFish(wg, flagFish, flagDog)
	go printDog(wg, flagDog, flagCat)

	wg.Wait()
	defer close(flagCat)

	fmt.Println("操作完成")

}

func printCat(wg *sync.WaitGroup, flagCat, flagFish chan struct{}) {
	var count int
	for {
		if count == 100 {
			wg.Done()
			return
		}
		<-flagCat
		fmt.Println("cat ", count+1, "次")
		count++
		flagFish <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, flagFish, flagDog chan struct{}) {
	var count int
	for {
		if count == 100 {
			wg.Done()
			return
		}
		<-flagFish
		fmt.Println("fish ", count+1, "次")
		count++
		flagDog <- struct{}{}
	}
}

func printDog(wg *sync.WaitGroup, flagDog, flagCat chan struct{}) {
	var count int
	for {
		if count == 100 {
			wg.Done()
			return
		}
		<-flagDog
		fmt.Println("dog ", count+1, "次")
		count++
		flagCat <- struct{}{}
	}
}

/**
参考材料
  1.使用协程，按顺序打印cat、dog、fish各100次：https://blog.csdn.net/qq_37102984/article/details/121791579

*/
