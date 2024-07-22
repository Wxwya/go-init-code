package model

import (
	"xwya/model/global"
)

type Field struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:字段ID"`
	FormID     uint   `json:"form_id" gorm:"not null;index;comment:表单的id"`
	FieldType  uint   `json:"field_type" gorm:"not null;comment:字段类型"`
	FieldName  string `json:"field_name" gorm:"type:varchar(30);not null;comment:字段名称"`
	FieldSize  string `json:"field_size" gorm:"type:varchar(10);comment:字段大小"`
	IsNull     *bool  `json:"is_null" gorm:"default:0;comment:是否为null"`
	IsPrimary  *bool  `json:"is_primary" gorm:";default:0;comment:是否为主键"`
	IsUnique   *bool  `json:"is_unique" gorm:"not null;default:0;comment:是否唯一值"`
	IsForeign  *bool  `json:"is_foreign" gorm:"not null;default:0;comment:当前字段是否为外键"`
	References string `json:"references" gorm:"comment:外键表对对应的字段"`
	IndexName  string `json:"index_name" gorm:"comment:索引名称"`
	IsValNull  *bool  `json:"is_val_null" gorm:"default:0;comment:为空时该字段是否返回"`
	global.Global
}
