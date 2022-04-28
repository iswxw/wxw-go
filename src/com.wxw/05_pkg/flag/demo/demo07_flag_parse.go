/*
@Time : 2022/4/28 17:18
@Author : weixiaowei
@File : demo07_flag_parse
*/
package main

import (
	"flag"
	"fmt"
	"time"
)

/**
  使用方法
     1. go run demo07_flag_parse.go -help  // 打印帮助信息
     2. 使用flag参数方法：./flag_demo -name 小明 --age 18 -married=false -d=1h30m
     3. 使用非flag参数的方法：./flag_demo a b c
*/

func main() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()

	fmt.Println(name, age, married, delay)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())

	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())

	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

// usage
//go run demo07_flag_parse.go -help
//Usage of /var/folders/vd/ms9cv4zj187g_wd6bzt4jq2r0000ks/T/go-build3947223647/b001/exe/demo07_flag_parse:
//-age int
//年龄 (default 18)
//-d duration
//延迟的时间间隔
//-married
//婚否
//-name string
//姓名 (default "张三")
