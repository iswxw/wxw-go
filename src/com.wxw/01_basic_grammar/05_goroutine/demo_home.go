/*
@Time: 2021/10/24 22:44
@Author: wxw
@File: case
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		go Add(i, i)
	}
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
