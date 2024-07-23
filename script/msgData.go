package script

import "xwya/model"

var DefaultMsgCodeData = []model.MsgCode{
	{Code: 200, Key: "Success", Msg: "Success"},                       // 成功
	{Code: 400, Key: "Error", Msg: "Error"},                           // 请求错误
	{Code: 401, Key: "Unauthorized", Msg: "Unauthorized"},             // 未授权
	{Code: 403, Key: "Forbidden", Msg: "Forbidden"},                   // 禁止访问
	{Code: 404, Key: "NotFound", Msg: "NotFound"},                     // 资源未找到
	{Code: 500, Key: "InternalError", Msg: "InternalError"},           // 服务器内部错误
	{Code: 503, Key: "ServiceUnavailable", Msg: "ServiceUnavailable"}, // 服务器当前无法处理请求
	{Code: 1000, Key: "TokenExist", Msg: "Token does not exist"},      // token不存在
	{Code: 1001, Key: "TokenExpired", Msg: "Token has expired"},       // token已过期
	{Code: 1002, Key: "TokenInvalid", Msg: "Token is invalid"},        // token无效
	{Code: 1003, Key: "TokenError", Msg: "Token error"},               // token错误
	{Code: 2000, Key: "UserNotExist", Msg: "User does not exist"},     // 用户不存在
	{Code: 2001, Key: "UserExist", Msg: "User already exists"},        // 用户已存在
	{Code: 2002, Key: "PasswordError", Msg: "Password error"},         // 密码错误
}
