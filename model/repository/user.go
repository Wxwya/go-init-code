package repository

import "xwya/model/global"

type QueryUserList struct {
	Username string `json:"username"`
	global.Page
}
