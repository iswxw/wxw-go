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
	"github.com/demo_utils/mylogger"
	"time"
)

// 测试我们自己写的日志库
func main() {

	// 写到控制台
	//funcConsoleLog()

	// 写日志到文件
	funcFileLog()
}

func funcConsoleLog() {
	log := mylogger.NewLog("INFO")
	for {
		log.Debug("这是一条Debug日志!!")
		log.Info("这是一条Info日志!!")
		log.Warning("这是一条Warning日志!!")
		id := 1000
		name := "理想"
		log.Error("这是一条Error日志!!id: %d,name: %s", id, name)
		log.Fatal("这是一条Fatal日志!!")
		time.Sleep(2 * time.Second)
		fmt.Println("————————————")
	}
}

func funcFileLog() {
	fmt.Println("文件日志开始写入...")
	// 10MB
	fileLogger := mylogger.NewFileLogger("INFO", "./doc/log", "wxw_go.log", 1024*1024*10)
	for {
		fileLogger.Debug("文件日志DEBUG级别")
		fileLogger.Info("文件日志INFO级别")
		fileLogger.Warning("文件日志WARNING级别")
		fileLogger.Error("文件日志ERROR级别")
		fileLogger.Fatal("文件日志FATAL级别")
		time.Sleep(2 * time.Second)
	}
}
