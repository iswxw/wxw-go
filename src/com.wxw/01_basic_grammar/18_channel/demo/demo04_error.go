/*
@Time: 2022/7/25 23:08
@Author: wxw
@File: demo04_error
*/
package main

import "fmt"

func main() {

	var name int
	// 定义一个int类型的通道
	c := make(chan int)
	c <- name
	fmt.Println(<-c)
}

// 输出结果
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan send]:
//main.main()