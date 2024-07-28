package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateField(c *gin.Context) {
	var data model.Field
	c.ShouldBindJSON(&data)
	if err := server.GenerateField(&data); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func DeleteField(c *gin.Context) {
	var ids map[string][]uint
	c.ShouldBindJSON(&ids)
	if err := server.DeleteField(ids["ids"]); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}

func GetFieldInfo(c *gin.Context) {
	id := c.Param("id")
	data, err := server.GetFieldInfo(id)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}

func GetFieldList(c *gin.Context) {
	var page entity.Page
	c.ShouldBindJSON(&page)
	list, total, err := server.GetFieldList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]interface{}{"list": list, "total": total})
}
