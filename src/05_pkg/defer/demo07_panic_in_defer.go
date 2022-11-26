/*
@Time: 2022/2/20 17:29
@Author: wxw
@File: demo07_panic_in_defer
*/
package main

import "fmt"

// defer中包含panic
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}
