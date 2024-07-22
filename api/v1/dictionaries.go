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
	if err := c.ShouldBindJSON(&dictionaries); err != nil {
		str := utils.TranslateValidationError(err)
		msgjson.ErrorValidateMsg(c, str)
		return
	}
	if err := server.GenerateDictionary(&dictionaries); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}
func DeleteDictionary(c *gin.Context) {
	var ids map[string][]int
	c.ShouldBindJSON(&ids)
	if err := server.DeleteDictionary(ids); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), nil)
}
func GetDictionaryList(c *gin.Context) {
	var page repository.QueryDictionary
	c.ShouldBindJSON(&page)
	data, total, err := server.GetDictionaryList(&page)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), map[string]interface{}{"data": data, "total": total})
}
func GetDictionaryInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		msgjson.ErrorMsg(c, msg.Error)
		return
	}
	data, err := server.GetDictionaryInfo(id)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, msg.GetMsg(msg.Success), data)
}
