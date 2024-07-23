package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initGlobalModelRouter(r *gin.RouterGroup) {
	globalModelRouter := r.Group("/lobalModel")
	{
		globalModelRouter.POST("generateGlobalModel", v1.GenerateGlobalModel)
		globalModelRouter.POST("deleteGlobalModel", v1.DeleteGlobalModel)
		globalModelRouter.POST("getGlobalModelList", v1.GetGlobalModelList)
		globalModelRouter.GET("getGlobalModelInfo", v1.GetGlobalModelInfo)
	}
}
