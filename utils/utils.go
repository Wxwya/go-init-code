package utils

import (
	"errors"
	"fmt"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
func ShouldBindJSON(c *gin.Context, data any) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		e, _ := err.(validator.ValidationErrors)
		HandleValidateError(c, e.Translate(trans))
		return false
	}
	return true
}
