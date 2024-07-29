package entity

type QueryProject struct {
	ProjectName string `json:"project_name" `
	SqlHost     string `json:"sql_host"`
	Frame       string `json:"frame"`
	Page
}
type QueryDictionary struct {
	DictType any `json:"dict"`
	Page
}

type QueryForm struct {
	// 必填
	ProjectID uint   `json:"project_id" binding:"required"`
	FormName  string `json:"form_name"`
	Page
}

type QueryMsgCode struct {
	MsgKey string `json:"msg_key"`
	Code   any    `json:"code"`
	Page
}

type QueryController struct {
	Version string `json:"version"`
	Page
}

type QueryControllerProcessor struct {
	Params   string `json:"params"`
	FuncName string `json:"func_name"`
	Page
}
type QueryUserList struct {
	Username string `json:"username"`
	Page
}
