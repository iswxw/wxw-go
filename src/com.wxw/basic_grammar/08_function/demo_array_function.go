/*
@Time: 2021/9/21 17:16
@Author: wxw
@File: demo_array_function 函数
@Link: https://www.runoob.com/go/go-functions.html
*/
package main

import "fmt"

func main() {
	/*数组长度为 5*/
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/*数组作为参数传递给函数*/
	avg = getAverage(balance, 5)

	/*输出返回的平均数*/
	fmt.Printf("平均值为：%f", avg)
}

// 定义一个函数
func getAverage /*函数名称*/ (arr [5]int, size int /*函数入参*/) float32 /*返回值类型*/ {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
