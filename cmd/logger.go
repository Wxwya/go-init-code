package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
	"xwya/utils"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	// 创建一个新的 logrus 实例
	utils.Logger = logrus.New()
	if _, err := os.Stat(utils.LogPath); os.IsNotExist(err) {
		err := os.Mkdir(utils.LogPath, os.ModePerm)
		if err != nil {
			fmt.Println("创建日志目录失败:", err)
		}
	}
	logTime := utils.LogPath + utils.LogName
	file, err := os.OpenFile(logTime, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		utils.Logger.SetOutput(file)
	}
	// 设置 logrus 日志级别
	utils.Logger.SetLevel(logrus.InfoLevel)

	// 设置 logrus 日志格式
	utils.Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
func Shutdown() {
	// 等待一段时间以确保所有日志都被刷新
	time.Sleep(2 * time.Second)
	// 关闭日志文件
	if file, ok := utils.Logger.Out.(*os.File); ok {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}
	log.Println("日志已关闭")
}
