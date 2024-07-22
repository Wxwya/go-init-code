package v1

import (
	"time"
	m "xwya/middleware"
	"xwya/model"
	"xwya/server"
	"xwya/utils"
	"xwya/utils/msg"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var info model.User
	c.ShouldBindJSON(&info)
	code, user := server.ValidateUser(info.Account, info.Password)
	if code != msg.Success {
		msgjson.ErrorMsg(c, code)
		return
	}
	token, code := m.SetToken(*user)
	if code != msg.Success {
		c.Set("err", "生成token失败")
		msgjson.ServerErrorMsg(c)
		return
	}
	c.SetCookie("xwya_admin_token", "Bearer "+token, int(time.Now().Add(time.Hour*72).Unix()), "/", "", false, true)
	msgjson.SuccessMsg(c, "ok", nil)
}

func BackloginOut(c *gin.Context) {
	c.SetCookie("xwya_admin_token", "", -1, "/", "", false, true)
	// c.JSON(http.StatusOK, gin.H{,"status": "ok"})
	msgjson.SuccessMsg(c, "退出成功", nil)
}

// 上传单文件
func UpFile(c *gin.Context) {
	// 跨域
	// origin := c.Request.Header.Get("Origin")
	// c.Header("Access-Control-Allow-Origin", origin)
	// 接收其他传参
	// directoryId, _ := strconv.Atoi(c.PostForm("directoryId"))

	file, _ := c.FormFile("file")

	fileName := utils.GenerateFileName(file.Filename)
	// fileSize := file.Size // 获取文件大小
	// fileType := file.Header.Get("Content-Type") // 获取文件类型
	newFileName := utils.FilePath + fileName
	err := c.SaveUploadedFile(file, newFileName)

	if err != nil {
		msgjson.ErrorMsg(c, msg.Error_UploadFile)
		return
	}
	msgjson.SuccessMsg(c, "上传成功", map[string]interface{}{
		"url": newFileName,
	})
}

// 上传多文件
func UpManyFile(c *gin.Context) {
	// origin := c.Request.Header.Get("Origin")
	// c.Header("Access-Control-Allow-Origin", origin)
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		fileName := utils.GenerateFileName(file.Filename)
		newFileName := utils.FilePath + fileName
		err := c.SaveUploadedFile(file, newFileName)
		if err != nil {
			msgjson.ErrorMsg(c, msg.Error_UploadFile)
			return
		}
	}
	msgjson.SuccessMsg(c, "上传成功", nil)

}
