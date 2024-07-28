package msg

const (
	Success     = 200 // 正确
	Error       = 400 // 错误
	NotFound    = 404 // 未找到
	ServerError = 500 // 服务器错误
	// code 1000开头 token相关
	TokenExist     = 1001 // token 错误
	TokenTimeout   = 1002 // token 过期
	TokenMalformed = 1003 // token 格式错误

	// code 2000开头全局相关
	NoID           = 2001 // 请传入id
	NoDeleteAdmin  = 2002 // 管理员不能删除
	PleaseTryAgain = 2005 // 请稍后在尝试
	// code 3000开头用户相关
	UserExist         = 3001 // 用户已存在
	UserNotExist      = 3002 // 用户不存在
	UserPasswordWrong = 3003 // 用户密码错误
	UserDisabled      = 3004 // 用户已被冻结
	// code 4000开头文件相关
	UploadFile = 4001
)

var codeMsg = map[int]string{
	Success:           "ok",
	Error:             "fail",
	NotFound:          "资源未找到",
	ServerError:       "服务器错误",
	TokenExist:        "token 不存在",
	TokenTimeout:      "token 已过期",
	TokenMalformed:    "token 格式错误",
	UserExist:         "用户已存在",
	UserNotExist:      "用户不存在",
	UserPasswordWrong: "用户密码错误",
	UploadFile:        "上传文件失败",
	NoID:              "请检查id",
	NoDeleteAdmin:     "不允许管理员用户删除",
	UserDisabled:      "用户已被冻结",
	PleaseTryAgain:    "请稍后在尝试",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
