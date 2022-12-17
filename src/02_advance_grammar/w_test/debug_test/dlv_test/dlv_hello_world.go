/*
@Time: 2022/3/12 22:47
@Author: wxw
@File: gdb_test
*/
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
