/*
@Time: 2022/4/9 22:26
@Author: wxw
@File: demo_mult_params
*/
package main

import (
	"fmt"
	"strings"
)

/**
  go 函数的可变参数
  1. 可变参数的使用场景
      - 避免创建仅作传入参数用的临时切片
      - 当参数数量未知
      - 传达你希望增加可读性的意图
  2. 相关资料
     - https://studygolang.com/articles/11965
     - https://www.pianshen.com/article/2474162160/
  3. 相关源码
     - func Prinln(a ...interface{})
  4. 功能
    - 混合使用可变参数及非可变参数
    - 对于函数式编程的实现
    - 可变配置模式
*/

func main() {
	fmt.Println(sum(1.1, 2.2, 3.3))
	fmt.Println("========")

	fmt.Println(ToIP(255))
	fmt.Println(ToIP(10, 1))
	fmt.Println(ToIP(127, 0, 0, 1))

	fmt.Println("======") // // output: "#01: carl sagan"
	fmt.Println(toFullName(1, "carl", "sagan"))
}

func sum(numbers ...float64) (res float64) {
	for _, number := range numbers {
		res += number
	}
	return
}

// 根据 parts 的长度返回一个字符串类型的 IP 地址，并且具有缺省值 —— 0。
func ToIP(parts ...byte) string {
	parts = append(parts, make([]byte, 4-len(parts))...)
	return fmt.Sprintf("%d.%d.%d.%d", parts[0], parts[1], parts[2], parts[3])
}

func toFullName(id int, names ...string) string {
	return fmt.Sprintf("#%02d: %s", id, strings.Join(names, " "))
}
