package v1

import (
	"strconv"
	"xwya/model"
	"xwya/model/global"
	"xwya/model/repository"
	"xwya/server"
	"xwya/utils/msgjson"

	"github.com/gin-gonic/gin"
)

func CreateAndUpdateUser(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	if err := server.CreateAndUpDateUser(&user); err != nil {
		c.Set("err", "创建用户或者更新用户失败"+err.Error())
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, "创建用户或者更新用户成功", nil)
}

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	info, err := server.GetUserInfo(user.(model.User).ID)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}

	msgjson.SuccessMsg(c, "获取用户信息成功", info)
}

func GetUserList(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	info := repository.QueryUserList{
		Page: global.Page{
			PageSize: pagesize,
			PageNum:  pagenum,
		},
		Username: username,
	}
	userList, total, err := server.GetUserList(&info)
	if err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, "获取用户列表成功", map[string]interface{}{"total": total, "userList": userList})
}
func DeleteUser(c *gin.Context) {
	ids := map[string][]int{}
	c.ShouldBind(&ids)
	if err := server.DeleteUser(ids["ids"]); err != nil {
		msgjson.ServerErrorMsg(c)
		return
	}
	msgjson.SuccessMsg(c, "删除用户成功", nil)
}
