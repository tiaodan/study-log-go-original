package main

import "study-json-go-origina/logger"

func main() {
	// 设置日志级别为debug
	logger.SetLogLevel(logger.LevelDebug)

	logger.Debug("这是一条调试信息")
	logger.Info("这是一条普通信息")
	logger.Warn("这是一条警告信息")
	logger.Error("这是一条错误信息")

	logger.Debug("这是一条调试信息-不带百分号", "debug")
	logger.Info("这是一条普通信息-不带百分号", "info")
	logger.Warn("这是一条警告信息-不带百分号", "warn")
	logger.Error("这是一条错误信息-不带百分号", "error")

	logger.Debug("这是一条调试信息-带百分号 %v", "debug")
	logger.Info("这是一条普通信息-带百分号 %v", "info")
	logger.Warn("这是一条警告信息-带百分号 %v", "warn")
	logger.Error("这是一条错误信息-带百分号 %v", "error")

	// 改变日志级别为Error
	logger.SetLogLevel(logger.LevelError)

	logger.Debug("这条调试信息不会显示") // 不会显示
	logger.Info("这条普通信息不会显示")  // 不会显示
	logger.Error("只有错误信息会显示")
}
