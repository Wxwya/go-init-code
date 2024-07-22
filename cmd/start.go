package cmd

import "xwya/router"

func Start() {
	defer CloseMysql()
	// defer CloseRedis()
	defer Shutdown()
	InitMysql()
	// InitRedis()
	InitLogger()
	InitValidate()

	router.InitRouter()

}
