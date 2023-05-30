// @Time : 2023/5/30 16:13
// @Author : xiaoweiwei
// @File : defer_exception_test

package _9_defer

import (
	"fmt"
	"testing"
)

func TestRecovery(t *testing.T) {
	divide(10, 0)
	fmt.Println("Program continues...")
}

// divide 当除数为 0 时，会触发 panic 异常，但由于使用了 defer 和 recover，程序不会直接崩溃
func divide(a, b int) {
	defer func() {
		if r := recover(); r != any(nil) {
			fmt.Println("Recovered form panic:", r)
		}
	}()
	result := a / b
	fmt.Println("Result:", result)
}
