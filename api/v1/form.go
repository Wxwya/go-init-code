package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateForm(c *gin.Context) {
	var form model.Form
	if flag := utils.ShouldBindJSON(c, &form); !flag {
		return
	}
	if err := server.GenerateForm(&form); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteForm(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteForm(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetFormList(c *gin.Context) {

	var page entity.QueryForm
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetFormList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetFormInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetFormInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}
