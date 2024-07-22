package msg

const (
	Success     = 200 // 正确
	Error       = 400 // 错误
	NotFound    = 404 // 未找到
	ServerError = 500 // 服务器错误
	// code 1000开头 token相关
	Error_TokenExist     = 1001 // token 错误
	Error_TokenTimeout   = 1002 // token 过期
	Error_TokenInvalid   = 1003 // token 无效
	Error_TokenMalformed = 1004 // token 格式错误

	// code 2000开头用户相关
	Error_UserExist         = 2001 // 用户已存在
	Error_UserNotExist      = 2002 // 用户不存在
	Error_UserPasswordWrong = 2003 // 用户密码错误
	// code 3000开头文件相关
	Error_UploadFile = 3001
)

var codeMsg = map[int]string{
	Success:                 "ok",
	Error:                   "fail",
	NotFound:                "资源未找到",
	ServerError:             "服务器错误",
	Error_TokenExist:        "token 不存在",
	Error_TokenTimeout:      "token 已过期",
	Error_TokenInvalid:      "token 不正确",
	Error_TokenMalformed:    "token 格式错误",
	Error_UserExist:         "用户已存在",
	Error_UserNotExist:      "用户不存在",
	Error_UserPasswordWrong: "用户密码错误",
	Error_UploadFile:        "上传文件失败",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
