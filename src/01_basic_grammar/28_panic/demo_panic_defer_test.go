/*
@Time: 2023/3/16 22:14
@Author: wxw
@File: demo_panic_defer
*/
package _8_panic

import (
	"fmt"
	"testing"
)

func TestPanicDefer(t *testing.T) {
	defer func() {
		fmt.Println("defer in main")
	}()
	f1()
	fmt.Println("Returned normally from f1.")
}

func f1() {
	/*defer func() {
	    if r := recover(); r != nil {
	        fmt.Println("Recovered in f", r)
	    }
	}()*/
	defer func() {
		fmt.Println("defer in f1")
	}()
	fmt.Println("Calling g1.")
	g1(0)
	fmt.Println("Returned normally from g1.")
}

func g1(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(any(fmt.Sprintf("%v", i)))
	}
	defer fmt.Println("Defer in g1", i)
	fmt.Println("Printing in g1", i)
	g1(i + 1)
}
