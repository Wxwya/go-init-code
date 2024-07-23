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

func GenerateController(c *gin.Context) {
	var controller model.Controller
	if flag := utils.ShouldBindJSON(c, &controller); !flag {
		return
	}
	if err := server.GenerateController(&controller); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func DeleteController(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteController(&ids); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func GetControllerList(c *gin.Context) {
	var page repository.QueryController
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetControllerList(&page)
	if err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetControllerInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		msgjson.HandleError(c, msg.Error)
		return
	}
	data, err := server.GetControllerInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		msgjson.HandleError(c, code)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), data)
}
