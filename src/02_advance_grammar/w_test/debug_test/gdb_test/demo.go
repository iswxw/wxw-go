// @Time : 2022/12/17 17:58
// @Author : xiaoweiwei
// @File : demo

package main

import "fmt"

func add(a, b int) int {
	sum := 0
	sum = a + b
	return sum
}

func main() {
	sum := add(10, 20)
	fmt.Println(sum)
}
