/*
@Time: 2022/2/20 17:31
@Author: wxw
@File: demo08_param_function
*/
package main

import "fmt"

func function(index int, value int) int {

	fmt.Println(index)

	return index
}

func main() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}
