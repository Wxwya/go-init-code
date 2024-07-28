package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initFieldRouter(r *gin.RouterGroup) {
	fieldRouter := r.Group("field")
	{
		fieldRouter.POST("generateField", v1.GenerateField)
		fieldRouter.POST("deleteField", v1.DeleteField)
		fieldRouter.GET("fieldinfo/:id", v1.GetFieldInfo)
		fieldRouter.POST("fieldlist", v1.GetFieldList)
	}
}
