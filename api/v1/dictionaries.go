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

func GenerateDictionary(c *gin.Context) {
	var dictionaries model.Dictionaries
	if flag := utils.ShouldBindJSON(c, &dictionaries); !flag {
		return
	}
	if err := server.GenerateDictionary(&dictionaries); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}
func DeleteDictionary(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)
	if err := server.DeleteDictionary(ids); err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), nil)
}
func GetDictionaryList(c *gin.Context) {
	var page repository.QueryDictionary
	if flag := utils.ShouldBindJSON(c, &page); !flag {
		return
	}
	data, total, err := server.GetDictionaryList(&page)
	if err != nil {
		msgjson.HandleServerError(c)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), map[string]interface{}{"data": data, "total": total})
}
func GetDictionaryInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		msgjson.HandleError(c, msg.Error)
		return
	}
	data, err := server.GetDictionaryInfo(id)
	if code := utils.IsErrRecordNotFound(err); code != msg.Success {
		msgjson.HandleError(c, code)
		return
	}
	msgjson.HandleSuccess(c, msg.GetMsg(msg.Success), data)
}
