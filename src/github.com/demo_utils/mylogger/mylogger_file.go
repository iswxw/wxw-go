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

// 根据指定文件路径和文件名 打开文件
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
		if f.checkSize(f.fileObj) {
			// 需要切割日志文件
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj) // 日志文件
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
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

// 日志分割
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前的日志文件完整路径
	// 拼接一个日志文件备份的名字
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr)
	// 1. 关闭当前的日志文件
	file.Close()
	// 2. rename 备份一下 xx.log -> xx.log.bak20200908
	os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	// 4. 将打开的新的日志文件 对象赋值给 f.fileObj
	return fileObj, nil
}

// 根据日志级别，判断是否需要记录日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// 检查文件大小
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	// 如果当前文件的大小 大于等于 日志文件的最大值，就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

// Debug 级别日志打印
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

// INFO
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

// WARNING
func (f *FileLogger) Warning(msg string) {
	f.log(WARNING, msg)
}

// ERROR
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}

// FATAL
func (f *FileLogger) Fatal(msg string) {
	f.log(FATAL, msg)
}

// 文件关闭
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
