package main

import (
	"io"
	"os"
	"study-json-go-origina/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个logrus实例
	log := logrus.New()

	// 设置自定义日志格式（控制台输出）
	log.SetFormatter(&logger.CustomFormatter{}) // 带颜色输出
	// log.SetFormatter(&CustomFileFormatter{}) // 不带颜色输出

	// 设置日志级别
	log.SetLevel(logrus.DebugLevel)

	// 创建一个文件用于写入日志
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
	}
	defer file.Close() // 关闭日志文件

	// 使用 io.MultiWriter 实现多写入器功能
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	log.Debug("这是一条调试信息11")
	log.Info("这是一条普通信息11")
	log.Warn("这是一条警告信息11")
	log.Error("这是一条错误信息11")

	log.Debug("这是一条调试信息-不带百分号", "debug")
	log.Info("这是一条普通信息-不带百分号", "info")
	log.Warn("这是一条警告信息-不带百分号", "warn")
	log.Error("这是一条错误信息-不带百分号", "error")

	log.Debugf("这是一条调试信息-带百分号 %v", "debug")
	log.Infof("这是一条普通信息-带百分号 %v", "info")
	log.Warnf("这是一条警告信息-带百分号 %v", "warn")
	log.Errorf("这是一条错误信息-带百分号 %v", "error")

	// 改变日志级别为Error
	log.Debug("--------------------改变日志级别为Error")
	log.SetLevel(logrus.ErrorLevel)

	log.Debug("这条调试信息不会显示") // 不会显示
	log.Info("这条普通信息不会显示")  // 不会显示
	log.Error("只有错误信息会显示")
}
