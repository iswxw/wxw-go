/*
@Time: 2022/3/12 22:47
@Author: wxw
@File: dbg_test
*/
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("Golang dbg test...")
	var argc = len(os.Args)
	var argv = append([]string{}, os.Args...)
	fmt.Printf("argc:%d\n", argc)
	fmt.Printf("argv:%v\n", argv)
	var var1 = 1
	var var2 = "golang dbg test"
	var var3 = []int{1, 2, 3}
	var var4 MyStruct
	var4.A = 1
	var4.B = "golang dbg my struct field B"
	var4.C = map[int]string{1: "value1", 2: "value2", 3: "value3"}
	var4.D = []string{"D1", "D2", "D3"}
	DBGTestRun(var1, var2, var3, var4)
	fmt.Println("Golang dbg test over")
}

type MyStruct struct {
	A int
	B string
	C map[int]string
	D []string
}

func DBGTestRun(var1 int, var2 string, var3 []int, var4 MyStruct) {
	fmt.Println("DBGTestRun Begin!\n")
	waiter := &sync.WaitGroup{}
	waiter.Add(1)
	go RunFunc1(var1, waiter)
	waiter.Add(1)
	go RunFunc2(var2, waiter)
	waiter.Add(1)
	go RunFunc3(&var3, waiter)
	waiter.Add(1)
	go RunFunc4(&var4, waiter)
	waiter.Wait()
	fmt.Println("DBGTestRun Finished!\n")
}
func RunFunc1(variable int, waiter *sync.WaitGroup) {
	fmt.Printf("var1:%v\n", variable)
	for {
		if variable != 123456 {
			continue
		} else {
			break
		}
	}
	time.Sleep(10 * time.Second)
	waiter.Done()
}
func RunFunc2(variable string, waiter *sync.WaitGroup) {
	fmt.Printf("var2:%v\n", variable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}
func RunFunc3(pVariable *[]int, waiter *sync.WaitGroup) {
	fmt.Printf("*pVar3:%v\n", *pVariable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}
func RunFunc4(pVariable *MyStruct, waiter *sync.WaitGroup) {
	fmt.Printf("*pVar4:%v\n", *pVariable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}
