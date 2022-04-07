/*
@Time : 2022/3/29 16:56
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := reflect.ValueOf(nil)
	fmt.Println(x.IsNil())
}

func f(arg interface{}) {
	i := arg.(int64)
	fmt.Println(i)
	switch v := arg.(type) {
	case float32, float64:
		fmt.Println(v == 0)
	}
}
