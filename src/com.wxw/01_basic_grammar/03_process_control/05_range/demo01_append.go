/*
@Time : 2022/3/29 17:59
@Author : weixiaowei
@File : demo01_append
*/
package main

import "fmt"

func main() {
	words := []string{"Go", "语言", "高性能", "编程"}
	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}
}
