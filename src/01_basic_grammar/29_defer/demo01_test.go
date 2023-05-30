// @Time : 2023/5/30 16:09
// @Author : xiaoweiwei
// @File : defer_test

package _9_defer

import (
	"fmt"
	"testing"
)

func TestExecuteOrder(t *testing.T) {
	defer fmt.Println("1. follow me") // 再执行
	defer fmt.Println("2. follow me") // 后执行
	fmt.Println("welcome me")         // 先执行
}
