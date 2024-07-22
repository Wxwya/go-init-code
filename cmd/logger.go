package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"
	"xwya/utils"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	// 创建一个新的 logrus 实例
	Logger = logrus.New()
	if _, err := os.Stat(utils.LogPath); os.IsNotExist(err) {
		err := os.Mkdir(utils.LogPath, os.ModePerm)
		if err != nil {
			fmt.Println("创建日志目录失败:", err)
		}
	}
	logTime := utils.LogPath + utils.LogName
	file, err := os.OpenFile(logTime, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	}
	// 设置 logrus 日志级别
	Logger.SetLevel(logrus.InfoLevel)

	// 设置 logrus 日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
func Shutdown() {
	// 等待一段时间以确保所有日志都被刷新
	time.Sleep(2 * time.Second)
	// 关闭日志文件
	if file, ok := Logger.Out.(*os.File); ok {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}
}

// 写日志的方法
func WriteLog(err error) {
	// 写日志的逻辑
	_, file, line, _ := runtime.Caller(1)
	Logger.WithField("file", file).WithField("line", line).Error(err.Error())
}
