/*
@Time : 2022/4/28 17:07
@Author : weixiaowei
@File : demo05_flag
*/
package main

import (
	"flag"
	"fmt"
)

// flag.Type(flag名, 默认值, 帮助信息)*Type
func main() {
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	fmt.Println(*name, *age, *married, *delay)
}

// 运行方式
// go run demo05_flag_type.go 张三 18 false 0
// 张三 18 false 0s
