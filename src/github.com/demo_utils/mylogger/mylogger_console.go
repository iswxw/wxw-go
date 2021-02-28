/*
 * @Time : 2021/2/28 10:34
 * @Author : wxw
 * @File : mylogger_console
 * @Software: GoLand
 * @Link:
 */
package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 往终端写日志相关内容

// 定义日志级别
type LogLevel uint16

// 定义全局常量
const (
	UNKNOWN LogLevel = iota // 0
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// 解析日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err) // 报错，打断程序运行。
	}
	return Logger{Level: level}
}

// 比较日志级别
func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

// Debug 级别日志打印
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		// 写入指定位置
		// fmt.Println(msg)

		// 输出时间
		now := time.Now()
		fmt.Printf("[%s] [Debug] %s \n", now.Format("2021-02-22 11:00:00"), msg)
	}
}

func (l Logger) Info(msg string) {
	// 写入指定位置
	if l.enable(INFO) {
		fmt.Println(msg)
	}
}

func (l Logger) Warning(msg string) {
	// 写入指定位置
	if l.enable(WARNING) {
		fmt.Println(msg)
	}
}

func (l Logger) Error(msg string) {
	// 写入指定位置
	if l.enable(ERROR) {
		fmt.Println(msg)
	}
}
func (l Logger) Fatal(msg string) {
	// 写入指定位置
	if l.enable(FATAL) {
		fmt.Println(msg)
	}
}
