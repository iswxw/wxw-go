/*
@Time: 2023/3/16 21:20
@Author: wxw
@File: demo_panic
*/
package _8_panic

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums[5])
}

func TestRecoverPanic(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	//defer func() {
	//    if r := recover(); r != nil {
	//        fmt.Println("Recovered in f", r)
	//    }
	//}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	fmt.Println("Printing in g", i)
	panic(i)
	fmt.Println("After panic in g", i)
}
