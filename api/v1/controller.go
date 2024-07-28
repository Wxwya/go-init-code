package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateController(c *gin.Context) {
	var controller model.Controller
	if flag := utils.ShouldBindJSON(c, &controller); !flag {
		return
	}
	if err := server.GenerateController(&controller); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteController(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteController(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetControllerList(c *gin.Context) {
	var page entity.QueryController
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetControllerList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetControllerInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetControllerInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}
