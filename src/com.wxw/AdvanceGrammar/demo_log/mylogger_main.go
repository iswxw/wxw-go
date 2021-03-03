/*
 * @Time : 2021/2/28 10:35
 * @Author : wxw
 * @File : mylogger_main
 * @Software: GoLand
 * @Link: https://www.bilibili.com/video/BV1FV411r7m8?p=82
 */
package main

import (
	"fmt"
	"github.com/demos/mylogger"
)

// 声明一个全局接口变量
var log mylogger.Logger

// 测试我们自己写的日志库
func main() {
	funcLog()
}

func funcLog() {
	fmt.Println("开始记录日志...")
	// log := mylogger.NewConsoleLogger("INFO") // 计入控制台
	log := mylogger.NewFileLogger("INFO", "./doc/log", "wxw_go.log", 1024*1024*10) // 10MB 计入文件
	for {
		log.Debug("这是一条Debug日志!!")
		log.Info("这是一条Info日志!!")
		log.Warning("这是一条Warning日志!!")
		id := 1000
		name := "理想"
		log.Error("这是一条Error日志!!id: %d,name: %s", id, name)
		log.Fatal("这是一条Fatal日志!!")
		//time.Sleep(2 * time.Second)
		fmt.Println("————————————")
	}
}
