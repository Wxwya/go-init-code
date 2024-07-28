package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateMsgCode(c *gin.Context) {
	var msgCode model.MsgCode
	if flag := utils.ShouldBindJSON(c, &msgCode); !flag {
		return
	}
	if err := server.GenerateMsgCode(&msgCode); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteMsgCode(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteMsgCode(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetMsgCodeList(c *gin.Context) {

	var page entity.QueryMsgCode
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetMsgCodeList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetMsgCodeInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetMsgCodeInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}
