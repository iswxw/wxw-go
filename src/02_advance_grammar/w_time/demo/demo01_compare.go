/*
@Time : 2022/1/25 15:52
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	getCurrentTime()

	// 2. Go 时间比较
	//timeCompare()
}

func timeCompare() {
	format := "2006-01-02 15:04:05"
	now := time.Now()

	//now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2015-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.After(a))
	fmt.Println(now.Before(a))
	fmt.Println(now.After(b))
	fmt.Println(now.Before(b))
	fmt.Println(a.After(b))
	fmt.Println(a.Before(b))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())
}

func getCurrentTime() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println(now.Format(format))

	fmt.Println(now.Second())

	// 年月日
	year := now.Year()
	month := now.Format("01")
	day := now.Day()
	fmt.Println(year, month, day)

	//或者
	year01 := now.Format("2006")
	month01 := now.Format("01")
	day01 := now.Format("02")
	hour := now.Format("15")
	min := now.Format("04")
	second := now.Format("05")
	fmt.Println(year01, month01, day01, hour, min, second)
}
