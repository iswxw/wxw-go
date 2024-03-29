package main

import (
	"errors"
	"fmt"
)

// 自定义一个函数，包括 加减法
type operator func(x, y int) int


// 函数类型作为参数 操作计算

func calculate(x, y int, op operator) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}


// 高阶函数
// 1. 接受其他的函数作为参数传入
func main() {

	// 加法计算
	op := func(x, y int) int {
		return x + y
	}
	v, err := calculate(1, 3, op)
	fmt.Println(v, err)
}
