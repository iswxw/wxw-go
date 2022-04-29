/*
@Time : 2022/4/29 10:42
@Author : weixiaowei
@File : 04_block
*/
package main

import "fmt"

// 如果可重名变量的类型不同,它们之间可能会存在“屏蔽”的现象
var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}
