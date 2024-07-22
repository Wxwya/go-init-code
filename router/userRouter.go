package router

import (
	v1 "xwya/api/v1"
	"xwya/middleware"

	"github.com/gin-gonic/gin"
)

func initUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user").Use(middleware.JwtToken())
	systemRouter := r.Group("/system")
	{
		userRouter.POST("adduser", v1.CreateAndUpdateUser)
		userRouter.GET("userinfo", v1.GetUserInfo)
		userRouter.GET("list", v1.GetUserList)
		userRouter.POST("delete", v1.DeleteUser)
	}
	{
		systemRouter.POST("/login", v1.Login)
		systemRouter.POST("/logout", v1.BackloginOut)
		systemRouter.POST("/upfile", middleware.JwtToken(), v1.UpFile)
		systemRouter.POST("upmanyfile", middleware.JwtToken(), v1.UpManyFile)
	}
}
