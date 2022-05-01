/*
@Time : 2022/4/29 11:17
@Author : weixiaowei
@File : demo_type
*/
package main

import "fmt"

func main() {

	// 1.int->string   能转成功不 //https://time.geekbang.org/column/article/13601
	s1 := string(rune(-1))
	s2 := string(-1)
	fmt.Println(s1, s2)

	// 2. string->[]byte

	s3 := string([]rune{'\u4F60', '\u597D'}) // 你好
	fmt.Println(s3)

}
