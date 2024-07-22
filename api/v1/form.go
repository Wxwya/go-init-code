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

func GenerateForm(c *gin.Context) {
	var project model.Form
	if err := c.ShouldBindJSON(&project); err != nil {
		str := utils.TranslateValidationError(err)
		msgjson.ErrorValidateMsg(c, str)
		return
	}
	if err := server.GenerateForm(&project); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}

func DeleteForm(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteForm(&ids); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}

func GetFormList(c *gin.Context) {

	var page repository.QueryForm
	if err := c.ShouldBindJSON(&page); err != nil {
		str := utils.TranslateValidationError(err)
		msgjson.ErrorValidateMsg(c, str)
		return
	}
	data, total, err := server.GetFormList(&page)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetFormInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		msgjson.ErrorMsg(c, msg.Error)
		return
	}
	data, err := server.GetFormInfo(id)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), data)
}
