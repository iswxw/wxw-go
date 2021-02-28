/*
 * @Time : 2021/2/28 10:33
 * @Author : wxw
 * @File : mylogger
 * @Software: GoLand
 * @Link:
 */
package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存路径
	fileName    string // 日志文件保存文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// 定义一个 NewFileLogger 构造函数，返回一个指针
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 按照文件名和文件路径 将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

//
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file Failed, err:%v \n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file Failed, err:%v \n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// file版 日志输出
func (f *FileLogger) log(lv LogLevel, format string, args ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n",
			// Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
			now.Format("2006-01-02 15:04:05"),
			getLogString(lv),
			fileName,
			funcName,
			lineNo,
			msg)
		if lv >= ERROR {
			// 如果要记录的日志大于等于 Error级别，则还要在日志文件中再记录一次
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s \n",
				// Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
				now.Format("2006-01-02 15:04:05"),
				getLogString(lv),
				fileName,
				funcName,
				lineNo,
				msg)
		}
	}
}

// 比较日志级别
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// Debug 级别日志打印
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

func (f *FileLogger) Warning(msg string) {
	f.log(WARNING, msg)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}
func (f *FileLogger) Fatal(msg string) {
	f.log(FATAL, msg)
}

// 文件关闭
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
