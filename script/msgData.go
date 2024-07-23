package script

import "xwya/model"

var DefaultMsgCodeData = []model.MsgCode{
	{Code: 200, MsgKey: "Success", Msg: "Success"},                       // 成功
	{Code: 400, MsgKey: "Error", Msg: "Error"},                           // 请求错误
	{Code: 401, MsgKey: "Unauthorized", Msg: "Unauthorized"},             // 未授权
	{Code: 403, MsgKey: "Forbidden", Msg: "Forbidden"},                   // 禁止访问
	{Code: 404, MsgKey: "NotFound", Msg: "NotFound"},                     // 资源未找到
	{Code: 500, MsgKey: "InternalError", Msg: "InternalError"},           // 服务器内部错误
	{Code: 503, MsgKey: "ServiceUnavailable", Msg: "ServiceUnavailable"}, // 服务器当前无法处理请求
	{Code: 1000, MsgKey: "TokenExist", Msg: "Token does not exist"},      // token不存在
	{Code: 1001, MsgKey: "TokenExpired", Msg: "Token has expired"},       // token已过期
	{Code: 1002, MsgKey: "TokenInvalid", Msg: "Token is invalid"},        // token无效
	{Code: 1003, MsgKey: "TokenError", Msg: "Token error"},               // token错误
	{Code: 2000, MsgKey: "UserNotExist", Msg: "User does not exist"},     // 用户不存在
	{Code: 2001, MsgKey: "UserExist", Msg: "User already exists"},        // 用户已存在
	{Code: 2002, MsgKey: "PasswordError", Msg: "Password error"},         // 密码错误
}
