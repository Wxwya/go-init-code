package router

import (
	"xwya/utils"

	"github.com/gin-gonic/gin"
)

func initStaticRouter(r *gin.Engine) {
	r.Static("/code", "./dist")
	r.Static("/static", "./"+utils.FilePath)
}
