package utils

import (
	"net/http"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data,omitempty"`
	Token string `json:"token,omitempty"`
}

func HandleResponse(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  msg.GetMsg(code),
		Data: data,
	})
}

func HandleValidateError(c *gin.Context, err validator.ValidationErrorsTranslations) {
	c.JSON(http.StatusBadRequest, &gin.H{
		"code": http.StatusBadRequest,
		"msg":  ParamError(err),
	})
}

func HandleServerError(c *gin.Context, err error) {
	errLog := map[string]any{
		"error_content": err.Error(),
		"status_code":   c.Writer.Status(),
		"path":          c.Request.RequestURI,
		"ip":            c.ClientIP(),
		"method":        c.Request.Method,
	}
	WriteLog(err, errLog)
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  "服务器错误",
	})
}

func HandleLoginResponse(c *gin.Context, token string) {
	c.JSON(http.StatusOK, Response{
		Code:  http.StatusOK,
		Msg:   "ok",
		Token: token,
	})
}
