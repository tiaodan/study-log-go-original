package main

import (
	"io"
	"os"
	"study-json-go-origina/log"

	"github.com/sirupsen/logrus"
)

func main() {
	// 2. 根据配置文件,设置日志相关,现在用logrus框架
	log.InitLog()

	// 获取日志实例
	log := log.GetLogger()

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
