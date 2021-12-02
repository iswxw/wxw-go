/*
@Time: 2021/9/21 17:31
@Author: wxw
@File: demo_mult_return_value
@Link: https://www.runoob.com/go/go-functions.html
       https://www.liwenzhou.com/posts/Go/07_pointer/
*/
package main

import "fmt"

func main() {
	a, b := swap("百度", "字节")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
