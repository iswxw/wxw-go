/*
@Time: 2023/8/25 9:54
@Author: wxw
@File: demo06_test
*/
package demo

import (
	"fmt"
	"testing"
)

func TestCache06(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}
