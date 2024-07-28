package utils

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func WriteLog(err error, args ...map[string]any) {
	_, file, line, _ := runtime.Caller(2)

	fmt.Println("错误:", err, "file", file, "line", line)
	Logger.WithField("file", file).WithField("line", line).Error(err.Error())
}
