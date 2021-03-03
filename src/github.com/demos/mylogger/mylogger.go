/*
 * @Time : 2021/2/28 17:26
 * @Author : wxw
 * @File : mylogger
 * @Software: GoLand
 * @Link:
 * @Vlog:
 */
package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 定义日志级别
type LogLevel uint16

type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

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

// 日志转字符串
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 运行时调用，获取调用者的 方法名、函数名、行数
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
