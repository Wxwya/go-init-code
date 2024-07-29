package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/script"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateProject(c *gin.Context) {
	var project model.Project
	if flag := utils.ShouldBindJSON(c, &project); !flag {
		return
	}
	if err := server.GenerateProject(&project); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteProject(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)
	if err := server.DeleteProject(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetProjectList(c *gin.Context) {
	var page entity.QueryProject
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetProjectList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetProjectInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetProjectInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}

func GenerateCode(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GenerateCode(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	script.InitProject(data)
	utils.HandleResponse(c, msg.Success, data)
}
