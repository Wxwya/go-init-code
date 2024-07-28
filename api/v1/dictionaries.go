package v1

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

func GenerateDictionary(c *gin.Context) {
	var dictionaries model.Dict
	if flag := utils.ShouldBindJSON(c, &dictionaries); !flag {
		return
	}
	if err := server.GenerateDictionary(&dictionaries); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}
func DeleteDictionary(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)
	if err := server.DeleteDictionary(ids); err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, nil)
}
func GetDictionaryList(c *gin.Context) {
	var page entity.QueryDictionary
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetDictionaryList(&page)
	if err != nil {
		utils.HandleServerError(c, err)
		return
	}
	utils.HandleResponse(c, msg.Success, map[string]any{"data": data, "total": total})
}
func GetDictionaryInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.HandleResponse(c, msg.Error, nil)
		return
	}
	data, err := server.GetDictionaryInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		utils.HandleResponse(c, code, nil)
		return
	}
	utils.HandleResponse(c, msg.Success, data)
}

func Cheshi(c *gin.Context) {
	data := server.Cheshi()
	utils.HandleResponse(c, msg.Success, data)
}
