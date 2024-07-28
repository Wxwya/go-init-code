package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	initStaticRouter(r)
	router := r.Group("/api")
	initDictionariesRouter(router)
	initProjectRouter(router)
	initFormRouter(router)
	initMsgCodeRouter(router)
	initControllerRouter(router)
	initControllerProcessorRouter(router)
	initGlobalModelRouter(router)
	initFieldRouter(router)
}
