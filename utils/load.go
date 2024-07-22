package utils

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	AppMode string
	Prot    string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisHostAndPort string
	RedisDb          string
	RedisPassword    string

	PassPrefix string

	LogPath string
	LogName string

	JwtKey string

	FilePath string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败", err)
	}
	LoadDatabase(file)   // 初始化数据库配置
	LoadServer(file)     // 初始化端口及环境配置
	LoadPassPrefix(file) // 初始化密码加密前缀
	LoadLogger(file)     // 初始化错误日志配置
	LoadJwtKey(file)     // 初始化jwtKey
	LoadFileInfo(file)   // 初始化Minio配置
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	// 读取AppMode区的内容
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	Prot = file.Section("server").Key("Port").MustString(":8080")
}
func LoadDatabase(file *ini.File) {
	// 读取dataBase区内容
	Db = file.Section("dataBase").Key("Db").MustString("mysql")
	DbHost = file.Section("dataBase").Key("DbHost").MustString("159.75.158.148")
	DbPort = file.Section("dataBase").Key("DbPort").MustString("3306")
	DbUser = file.Section("dataBase").Key("DbUser").MustString("root")
	DbPassword = file.Section("dataBase").Key("DbPassword").MustString("w73748196")
	DbName = file.Section("dataBase").Key("DbName").MustString("cheshi")
}

func LoadPassPrefix(file *ini.File) {
	PassPrefix = file.Section("md5").Key("PassPrefix").MustString("xwya")
}

func LoadLogger(file *ini.File) {
	LogPath = file.Section("log").Key("LogPath").MustString("logs/")
	LogName = file.Section("log").Key("LogName").MustString("error.log")
}
func LoadRedis(file *ini.File) {
	RedisHostAndPort = file.Section("redis").Key("RedisHostAndPort").MustString("159.75.158.148:6379")
	RedisDb = file.Section("redis").Key("RedisDb").MustString("0")
	RedisPassword = file.Section("redis").Key("RedisPassword").MustString("w73748196")
}
func LoadJwtKey(file *ini.File) {
	JwtKey = file.Section("jwt").Key("JwtKey").MustString("e9a62e0d-d520-465f-8188-c3e7b13977bd")
}
func LoadFileInfo(file *ini.File) {
	FilePath = file.Section("file").Key("filePath").MustString("static/")
}

func GenerateFileName(originalFileName string) string {
	timeStamp := time.Now().Unix()
	return fmt.Sprintf("%d_%s", timeStamp, originalFileName)
}
