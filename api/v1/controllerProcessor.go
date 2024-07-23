package v1

import (
	"fmt"
	"xwya/model"
	"xwya/model/repository"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
)

func GenerateControllerProcessor(c *gin.Context) {
	var controllerProcessor model.ControllerProcessor
	if flag := utils.ShouldBindJSON(c, &controllerProcessor); !flag {
		return
	}

	if err := server.GenerateControllerProcessor(&controllerProcessor); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func DeleteControllerProcessor(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteControllerProcessor(&ids); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}

func GetControllerProcessorList(c *gin.Context) {

	var page repository.QueryControllerProcessor
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	fmt.Println("controllerProcessor:@@@@@@", page)

	data, total, err := server.GetControllerProcessorList(&page)
	if err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), map[string]interface{}{
		"list":  data,
		"total": total,
	})
}

func GetControllerProcessorInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		msgjson.HandleError(c, msg.Error)
		return
	}
	data, err := server.GetControllerProcessorInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		msgjson.HandleError(c, code)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), data)
}
