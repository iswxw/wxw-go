/*
 * @Time : 2021/2/28 10:35
 * @Author : wxw
 * @File : mylogger_main
 * @Software: GoLand
 * @Link:
 */
package main

import (
	"fmt"
	"github.com/demo_utils/mylogger"
	"time"
)

/** 日志
https://www.bilibili.com/video/BV1FV411r7m8?p=82
*/

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("Error")
	for {
		log.Debug("这是一条Debug日志!!")
		log.Info("这是一条Info日志!!")
		log.Warning("这是一条Warning日志!!")
		log.Error("这是一条Error日志!!")
		log.Fatal("这是一条Fatal日志!!")
		time.Sleep(time.Second)
		fmt.Println("————————————")
	}

}
