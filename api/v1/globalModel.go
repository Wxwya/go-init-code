package v1

import (
	"xwya/model"
	"xwya/model/global"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
)

func GenerateGlobalModel(c *gin.Context) {
	var lobalModel model.GlobalModel
	if flag := utils.ShouldBindJSON(c, &lobalModel); !flag {
		return
	}
	if err := server.GenerateGlobalModel(&lobalModel); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func DeleteGlobalModel(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteGlobalModel(&ids); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func GetGlobalModelList(c *gin.Context) {

	var page global.Page
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetGlobalModelList(&page)
	if err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetGlobalModelInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		msgjson.HandleError(c, msg.Error)
		return
	}
	data, err := server.GetGlobalModelInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		msgjson.HandleError(c, code)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), data)
}
