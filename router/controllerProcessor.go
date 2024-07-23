package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initControllerProcessorRouter(r *gin.RouterGroup) {
	controllerProcessorRouter := r.Group("/controllerProcessor")
	{
		controllerProcessorRouter.POST("generateControllerProcessor", v1.GenerateControllerProcessor)
		controllerProcessorRouter.POST("deleteControllerProcessor", v1.DeleteControllerProcessor)
		controllerProcessorRouter.POST("getControllerProcessorList", v1.GetControllerProcessorList)
		controllerProcessorRouter.GET("getControllerProcessorInfo", v1.GetControllerProcessorInfo)
	}
}
