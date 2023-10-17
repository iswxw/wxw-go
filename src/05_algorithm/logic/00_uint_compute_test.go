// @Time : 2023/5/24 19:38
// @Author : xiaoweiwei
// @File : main_test

package logic

import (
	"fmt"
	"testing"
)

// 计算数字异常
func TestName(t *testing.T) {
	a := uint(1)
	b := uint(2)
	fmt.Println(a - b)
}

// 相关材料
// 1. Go无符号整数运算时反转问题:https://blog.csdn.net/swan_tang/article/details/122111847
