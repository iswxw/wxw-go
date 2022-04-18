/*
@Time : 2022/4/18 16:02
@Author : weixiaowei
@File : demo_struct_copy
*/
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	A int
	B int
	C string
	E string
}

type B struct {
	B int
	C string
	D int
	E string
}

// https://zhuanlan.zhihu.com/p/27762748
func main() {
	test1()
}

func test1() {
	a := &A{1, 1, "a", "b"}
	aj, _ := json.Marshal(a)
	b := new(B)
	_ = json.Unmarshal(aj, b)

	fmt.Printf("%+v", b)
}
