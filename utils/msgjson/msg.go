package msgjson

import (
	"net/http"
	"xwya/utils/msg"

	"github.com/gin-gonic/gin"
)

// Response is a common response structure
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Token string      `json:"token,omitempty"`
}

func SuccessMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}

func ErrorMsg(c *gin.Context, err int) {
	c.JSON(http.StatusOK, Response{
		Code: err,
		Msg:  msg.GetMsg(err),
	})
}

func ErrorValidateMsg(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  err,
	})
}

func ServerErrorMsg(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:  "服务器错误",
	})
}

func LoginMsg(c *gin.Context, msg string, token string) {
	c.JSON(http.StatusOK, Response{
		Code:  http.StatusOK,
		Msg:   msg,
		Token: token,
	})
}
