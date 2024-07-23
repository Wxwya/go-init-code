package model

import (
	"xwya/model/global"
	"xwya/utils/sqltype"
)

/*
	控制器中处理的方法

*/

type ControllerProcessor struct {
	ID              uint          `json:"id" gorm:"primaryKey;autoIncrement;comment:表ID"`
	ControllerID    uint          `json:"controller_id" gorm:"comment:所属控制器ID"`
	FuncName        string        `json:"func_name" gorm:"index:idx_func_name_method_id size:50;unique;not null;comment:方法名称"`
	FuncDescription string        `json:"func_description" gorm:"type:text;comment:方法描述"`
	FuncCode        string        `json:"func_code" gorm:"type:text;comment:方法代码"`
	MethodID        uint          `json:"method_id,omitempty" gorm:"index:idx_func_name_method_id;not null;comment:请求方法ID(对应字典表5)"`
	Method          string        `json:"method"  gorm:"size:10;not null;comment:请求方法"`
	JoinMethod      string        `json:"join_method" gorm:"type:enum('query', 'param');comment:Get请求传参方式(query||param)"`
	Params          sqltype.Array `json:"params" gorm:"comment:Get请求传入的参数"`
	SuccessCodeID   uint          `json:"success_code_id" gorm:"comment:成功状态码的id"`
	ErrorCodeID     uint          `json:"error_code_id" gorm:"comment:错误状态码的id"`
	BodyID          uint          `json:"body_id" gorm:"comment:post|delete|put等方法请求体(对应字典表的go类型3)"`
	global.Global
}
