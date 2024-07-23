package utils

import (
	"errors"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsErrRecordNotFound(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return msg.NotFound
	}
	if err != nil {
		return msg.ServerError
	}
	return msg.Success
}
func ShouldBindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {

		str := TranslateValidationError(err)
		msgjson.HandleValidateError(c, str)
		return false
	}
	return true
}
