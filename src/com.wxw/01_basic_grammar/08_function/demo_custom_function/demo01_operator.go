/*
@Time : 2022/4/29 20:48
@Author : weixiaowei
@File : demo01_operator
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	op := func(x, y int) int { return x + y }
	value, err := calculate(3, 6, op)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}
