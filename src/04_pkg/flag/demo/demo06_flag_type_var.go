/*
@Time : 2022/4/28 17:16
@Author : weixiaowei
@File : demo06_flag_type_var
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
func main() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	fmt.Println("输出结果：", name, age, married, delay)
}

// go run demo06_flag_type_var.go
// 输出结果： 张三 18 false 0s
