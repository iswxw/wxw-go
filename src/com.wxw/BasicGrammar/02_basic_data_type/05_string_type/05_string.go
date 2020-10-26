/*
@Time : 2020/10/26 19:58
@Author : wxw
@File : 05_string
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)

// 字符串
func main() {
	// \ 本来具有特殊含义，所以加 \ 表示单纯的 \ 不再解析
	path := "F:\\Study_GO\\Study_Project\\go\\wxw-go\\src\\com.wxw\\BasicGrammar\\02_basic_data_type\\05_string_type\\05_string.go"
	fmt.Printf("%s \n", path)

	s := "I‘m ok"
	fmt.Printf(s)

	// 多行字符串
	s2 := `
     年少无为、你看鲲鹏！
    `
	fmt.Printf(s2)

	// 字符串相关操作
	fmt.Println(len(path))

	// 字符串拼接
	name := "理想"
	age := "年轻"
	ss := name + age
	fmt.Println(ss)
	ssl := fmt.Sprintf("%s %s", name, age)
	// ssl := fmt.printf("%s %s", name, age)
	fmt.Println(ssl)

	// 字符串分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 字符串包含
	fmt.Println(strings.Contains(ss, "理想"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	// 子串出现的位置
	f4 := "abcd"
	fmt.Println(strings.Index(f4, "c"))

	// 拼接
	fmt.Println(strings.Join(ret, "_"))

}
