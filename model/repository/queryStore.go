package repository

import "xwya/model/global"

type QueryProject struct {
	ProjectName string `json:"project_name" `
	SqlHost     string `json:"sql_host"`
	Frame       string `json:"frame"`
	global.Page
}
type QueryDictionary struct {
	Code uint `json:"code"`
	global.Page
}

type QueryForm struct {
	// 必填
	ProjectID uint   `json:"project_id" binding:"required"`
	FormName  string `json:"form_name"`
	global.Page
}
