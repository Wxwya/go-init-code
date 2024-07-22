package v1

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
)

func GenerateProject(c *gin.Context) {
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		str := utils.TranslateValidationError(err)
		msgjson.ErrorValidateMsg(c, str)
		return
	}
	if err := server.GenerateProject(&project); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}

func DeleteProject(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteProject(&ids); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}

func GetProjectList(c *gin.Context) {

	var page repository.QueryProject
	c.ShouldBindJSON(&page)
	data, total, err := server.GetProjectList(&page)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetProjectInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		msgjson.ErrorMsg(c, msg.Error)
		return
	}
	data, err := server.GetProjectInfo(id)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), data)
}
