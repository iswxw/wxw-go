/*
 * @Time : 2021/2/28 16:18
 * @Author : wxw
 * @File : demo_time
 * @Software: GoLand
 * @Link: https://www.liwenzhou.com/posts/Go/go_time/
 * @vlog:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间格式化
	fmt.Println("-----时间格式化---------")
	formatDemo()
	fmt.Println("-----字符串向时间转换-----")
	// 字符串向时间转换
	StringDemo()

	fmt.Printf("当前时间：%s \n", time.Now())
	fmt.Printf("1970.1.1 年到现在的秒数：%v\n", time.Now().Unix())
	fmt.Printf("1970.1.1 到现在的纳秒数：%v\n", time.Now().UnixNano())

	// 基本用法
	timeDemo()

	// 定时器
	//tickDemo()
}

// 字符串向时间转换
func StringDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

// 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

// 基本用法
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
