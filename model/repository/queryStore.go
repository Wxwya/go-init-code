package repository

import "xwya/model/global"

type QueryProject struct {
	ProjectName string `json:"project_name" `
	SqlHost     string `json:"sql_host"`
	Frame       string `json:"frame"`
	global.Page
}
type QueryDictionary struct {
	Code interface{} `json:"code"`
	global.Page
}

type QueryForm struct {
	// 必填
	ProjectID uint   `json:"project_id" binding:"required"`
	FormName  string `json:"form_name"`
	global.Page
}

type QueryMsgCode struct {
	MsgKey string      `json:"msg_key"`
	Code   interface{} `json:"code"`
	global.Page
}

type QueryController struct {
	Version string `json:"version"`
	global.Page
}

type QueryControllerProcessor struct {
	Params   string `json:"params"`
	FuncName string `json:"func_name"`
	IsHooks  *bool  `json:"is_hooks"`
	global.Page
}
