package cmd

import (
	"fmt"
	"time"
	"xwya/model"
	"xwya/server"
	"xwya/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 连接数据库
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置为 true 时使用单数表名形式
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败请检查:", err)
		panic(err)
	}
	// Logger: logger.Default.LogMode(logger.Silent)}) //  或 logger.Silent、logger.Warn 等级
	database.Session(&gorm.Session{
		// 启用 GORM 的日志记录
		// 启用 GORM 的日志记录
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Println("连接数据库成功")
	server.Db = database

	// 设置连接池参数
	sqlDB, err := server.Db.DB()
	if err != nil {
		fmt.Println("连接数据库失败请检查:", err)
		panic(err)
	}
	// 最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最大连接时间
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	// 自动迁移
	migrate()

}

// 数据库迁移
func migrate() {
	//

	// &model.Repository{},
	// &model.Form{},
	// &model.Field{},
	// &model.GlobalModel{},

	// &model.GlobalModel{},
	// &model.Repository{},

	server.Db.AutoMigrate(
		&model.Dict{},
		&model.Project{},
		&model.GlobalModel{},
		&model.Entity{},
		&model.Form{},
		&model.Field{},
		&model.MsgCode{},
		&model.Controller{},
		&model.ControllerProcessor{},
	)

}

func CloseMysql() {
	sqlDB, err := server.Db.DB()
	if err != nil {
		fmt.Println("关闭数据库连接失败:", err)
		return
	}
	sqlDB.Close()
	fmt.Println("数据库连接已关闭")
}
