// 封装go自带log库, 分 debug、info、warn、error级别日志
package logger

import (
	"log"
	"os"
)

// 定义日志级别
type LogLevel int

// 常量, 具体日志级别
const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

// 变量
var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	logLevel    LogLevel = LevelInfo // 默认日志级别为info
)

// 初始化, 设置日志，都打印到文件里
func init() {
	logFile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(logFile) // 设置日志输出到文件
	debugLogger = log.New(log.Writer(), "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(log.Writer(), "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(log.Writer(), "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(log.Writer(), "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// SetLogLevel 设置日志级别
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// 打印debug级别日志
func Debug(format string, v ...interface{}) {
	if logLevel <= LevelDebug {
		debugLogger.Printf(format, v...)
	}
}

// 打印info级别日志
func Info(format string, v ...interface{}) {
	if logLevel <= LevelInfo {
		infoLogger.Printf(format, v...)
	}
}

// 打印warn级别日志
func Warn(format string, v ...interface{}) {
	if logLevel <= LevelWarn {
		warnLogger.Printf(format, v...)
	}
}

// 打印error级别日志
func Error(format string, v ...interface{}) {
	if logLevel <= LevelError {
		errorLogger.Printf(format, v...)
	}
}
