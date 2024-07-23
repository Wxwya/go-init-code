package model

import (
	"xwya/model/global"
	"xwya/utils/sqltype"
)

/*
存储库
在字典中生成查询的类型
*/
type Repository struct {
	ID              uint          `json:"id" gorm:"primaryKey;autoIncrement;comment:自增id"`
	Code            uint          `json:"code" gorm:"default:3"`
	FormID          uint          `json:"form_id" gorm:"comment:表单id"`
	FormField       sqltype.Array `json:"form_field" gorm:"not null;comment:表单里的字段ID"`
	GlobalModelID   uint          `json:"global_model_id" gorm:"comment:引用全局模型"`
	GlobalModelName string        `json:"global_model_name,omitempty" gorm:"comment:引用全局模型名称"`
	global.Global
}