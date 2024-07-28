package v1

import (
	"fmt"
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateControllerProcessor(c *gin.Context) {
	var controllerProcessor model.ControllerProcessor
	if flag := utils.ShouldBindJSON(c, &controllerProcessor); !flag {
		return
	}

	if err := server.GenerateControllerProcessor(&controllerProcessor); err != nil {
		utils.HandleServerError(c, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteControllerProcessor(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)

	if err := server.DeleteControllerProcessor(&ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetControllerProcessorList(c *gin.Context) {

	var page entity.QueryControllerProcessor
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	fmt.Println("controllerProcessor:@@@@@@", page)

	data, total, err := server.GetControllerProcessorList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{
		"list":  data,
		"total": total,
	})
}

func GetControllerProcessorInfo(c *gin.Context) {
	// 获取地址上的值
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetControllerProcessorInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}
