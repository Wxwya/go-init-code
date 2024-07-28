package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateGlobalModel(c *gin.Context) {
	var lobalModel model.GlobalModel
	if flag := utils.ShouldBindJSON(c, &lobalModel); !flag {
		return
	}
	if err := server.GenerateGlobalModel(&lobalModel); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteGlobalModel(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteGlobalModel(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetGlobalModelList(c *gin.Context) {

	var page entity.Page
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetGlobalModelList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetGlobalModelInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetGlobalModelInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}
