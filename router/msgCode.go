package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initMsgCodeRouter(r *gin.RouterGroup) {
	msgCodeRouter := r.Group("/msgCode")
	{
		msgCodeRouter.POST("generateMsgCode", v1.GenerateMsgCode)
		msgCodeRouter.POST("deleteMsgCode", v1.DeleteMsgCode)
		msgCodeRouter.POST("getMsgCodeList", v1.GetMsgCodeList)
		msgCodeRouter.GET("getMsgCodeInfo", v1.GetMsgCodeInfo)
	}
}
