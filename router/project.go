package router

import (
	v1 "xwya/api/v1"

	"github.com/gin-gonic/gin"
)

func initProjectRouter(r *gin.RouterGroup) {
	projectRouter := r.Group("/project")
	{
		projectRouter.POST("generateProject", v1.GenerateProject)
		projectRouter.POST("deleteProject", v1.DeleteProject)
		projectRouter.POST("getProjectList", v1.GetProjectList)
		projectRouter.GET("getProjectInfo", v1.GetProjectInfo)
	}
}
