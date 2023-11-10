/*
@Time : 2022/4/29 20:48
@Author : weixiaowei
@File : demo01_operator
*/
package demo_custom_function

import (
	"errors"
	"fmt"
	"testing"
)

type operate func(x, y int) int

type calculatorFunc func(x, y int) (int, error)

func genCalculator(op operate) calculatorFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New(fmt.Sprintf("Invalid operation"))
		}
		return op(x, y), nil
	}
}

// 高阶函数
// 1. 把其他的函数作为结果返回
func TestReturnResult(t *testing.T) {

	// 定义一个加法操作
	op := func(x, y int) int {
		return x + y
	}

	// 生成计算函数
	add := genCalculator(op)
	result, err := add(1, 2)

	fmt.Printf("The result: %d \nerror: %v \n", result, err)

}
