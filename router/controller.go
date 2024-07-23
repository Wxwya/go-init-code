package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initControllerRouter(r *gin.RouterGroup) {
	controllerRouter := r.Group("/controller")
	{
		controllerRouter.POST("generateController", v1.GenerateController)
		controllerRouter.POST("deleteController", v1.DeleteController)
		controllerRouter.POST("getControllerList", v1.GetControllerList)
		controllerRouter.GET("getControllerInfo", v1.GetControllerInfo)
	}
}
