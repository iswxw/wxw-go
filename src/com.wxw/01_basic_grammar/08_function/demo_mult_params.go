/*
@Time: 2022/4/9 22:26
@Author: wxw
@File: demo_mult_params
*/
package main

import "strings"

/**
  go 函数的可变参数
  1. 可变参数的使用场景
      - 避免创建仅作传入参数用的临时切片
      - 当参数数量未知
      - 传达你希望增加可读性的意图
  2. 相关资料
     - https://studygolang.com/articles/11965
  3. 相关源码
     - func Prinln(a ...interface{})
  4. 功能
    - 混合使用可变参数及非可变参数
    - 对于函数式编程的实现
    - 可变配置模式
*/
func main() {

	// 简单使用
	test01()
}

func test01() {
	toFullName("carl", "sagan")

	// output: "carl sagan"

	toFullName("carl")

	// output: "carl"

	toFullName()

	// output: ""
}

func toFullName(names ...string) string {
	return strings.Join(names, " ")
}
