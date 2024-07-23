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

func GenerateMsgCode(c *gin.Context) {
	var msgCode model.MsgCode
	if flag := utils.ShouldBindJSON(c, &msgCode); !flag {
		return
	}
	if err := server.GenerateMsgCode(&msgCode); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func DeleteMsgCode(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteMsgCode(&ids); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func GetMsgCodeList(c *gin.Context) {

	var page repository.QueryMsgCode
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetMsgCodeList(&page)
	if err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetMsgCodeInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		msgjson.HandleError(c, msg.Error)
		return
	}
	data, err := server.GetMsgCodeInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		msgjson.HandleError(c, code)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), data)
}
