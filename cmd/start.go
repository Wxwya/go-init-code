package cmd

func Start() {
	defer CloseMysql()
	// defer CloseRedis()
	defer Shutdown()
	InitMysql()
	// InitRedis()
	InitLogger()
	InitValidate()
	InitHttp()

}
