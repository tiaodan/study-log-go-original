# v0.0.0.1 2021-03-30
- 第一次提交, 实现日志分级打印，但只能打印到日志文件，不能在控制台显示

# v0.0.0.1 2021-03-30-v2
- 日志可以打印到控制台+文件

# v0.0.0.2
- 修改日志打印,可以不带%v.修改: logger.Error("创建失败:", result.Error) 会警告,只能带上%v才对. logger.Error("创建失败:", result.Error)

# v0.0.0.3
- 整合日志打印,looger.Debug() 参数带不带%v都可以

# v0.0.0.4
- 解决logger.Debug("xxxxx") 不打印

# v0.0.0.5
- 自己封装总是有点问题, 改用logrus框架
- 但是无法控制 输出到控制台(带颜色), 输出到文件(不带颜色)
- 能用配置文件,控制打印带颜色,还是不带颜色的
- 打印带占位符%v 就用log.Debugf(); 不带占位符用log.Debug()

# v0.0.0.6
- 改用logrus框架
- 把logger包改成log包
- 用了自定义logrus,就不建议用go 自带的log库了,因为可能会冲突,如果非要打印,用fmt.Println()
- 把logger.Debug()打印, 全换成log.Debugf()
