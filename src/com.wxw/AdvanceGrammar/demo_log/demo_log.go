/*
@Time : 2021/2/28 10:00
@Author : wxw
@File : demo_utils
@Software: GoLand
@link: https://www.liwenzhou.com/posts/Go/go_log/
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 定义日志输出IO
	log.Println("这是一条普通的日志！")
	// 配置
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) // 日志输出详细信息参数配置
	log.SetPrefix("[小W]")                                       // 日志前缀

	// 打印日志
	v := "就是很普通"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}

// 定义日志输出位置
func init() {
	logFile, err := os.OpenFile("/demo_utils.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file Field, err:", err)
		return
	}
	log.SetOutput(logFile)   // 写入文件
	log.SetOutput(os.Stdout) // 往终端写
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
