/*
 * @Time : 2021/2/28 10:34
 * @Author : wxw
 * @File : mylogger_console
 * @Software: GoLand
 * @Link:
 */
package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err) // 报错，打断程序运行。
	}
	return ConsoleLogger{Level: level}
}

// 比较日志级别
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

// 统一处理打印
func (c ConsoleLogger) log(lv LogLevel, format string, args ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n",
			// Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
			now.Format("2006-01-02 15:04:05"),
			getLogString(lv),
			fileName,
			funcName,
			lineNo,
			msg)
	}
}

// Debug 级别日志打印
func (c ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DEBUG, format, args...)
}

func (c ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(INFO, format, args...)
}

func (c ConsoleLogger) Warning(msg string) {
	c.log(WARNING, msg)
}

func (c ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ERROR, format, args...)
}

func (c ConsoleLogger) Fatal(msg string) {
	c.log(FATAL, msg)
}
