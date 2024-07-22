package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initFormRouter(r *gin.RouterGroup) {
	formRouter := r.Group("/form")
	{
		formRouter.POST("generateForm", v1.GenerateForm)
		formRouter.POST("deleteForm", v1.DeleteForm)
		formRouter.POST("getFormList", v1.GetFormList)
		formRouter.GET("getFormInfo", v1.GetFormInfo)
	}
}
