package model

import (
	"xwya/model/global"
)

/*
		References 字段需要设置当前表中的字段
	  ForeignKey 设置模型表中的字段
*/
type Field struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:字段ID"`
	FormID        uint   `json:"form_id,omitempty" gorm:"default:0;index;comment:表单的id"`
	FieldType     uint   `json:"field_type" gorm:"not null;comment:字段类型"`
	FieldName     string `json:"field_name" gorm:"type:varchar(30);not null;comment:字段名称"`
	FieldSize     uint   `json:"field_size" gorm:"comment:字段大小"`
	IsNull        *bool  `json:"is_null,omitempty" gorm:"default:0;comment:是否为null"`
	IsPrimary     *bool  `json:"is_primary,omitempty" gorm:";default:0;comment:是否为主键"`
	IsUnique      *bool  `json:"is_unique,omitempty" gorm:"not null;default:0;comment:是否唯一值"`
	ForeignKey    *bool  `json:"foreign_key,omitempty" gorm:"not null;default:0;comment:设置外键"`
	References    string `json:"references,omitempty" gorm:"comment:设置映射"`
	IndexName     string `json:"index_name,omitempty" gorm:"comment:索引名称(为复合索引使用)"`
	IsValNull     *bool  `json:"is_val_null,omitempty" gorm:"default:0;comment:为空时该字段是否返回"`
	GlobalModelID uint   `json:"global_model_id,omitempty" gorm:"index;comment:全局模型ID"`
	IsQuery       *bool  `json:"is_query,omitempty" gorm:"default:false;comment:是否是查询字段"`
	global.Global
}
