package model

import "xwya/entity"

/*
	全局模型
	创建表的时候需要自动生成 page结构体(页码,条数) 和 Global结构体(包含创建时间,自动更新时间)

*/
type GlobalModel struct {
	ID              uint     `json:"id" gorm:"primaryKey;autoIncrement;comment:全局模型的id"`
	GlobalModelName string   `json:"global_model_name" gorm:"size:50;not null;comment:全局模型的名称"`
	Fields          *[]Field `json:"fields,omitempty" gorm:"foreignKey:GlobalModelID;constraint:OnDelete:CASCADE;comment:一对多关系"`
	Entity          Entity   `json:"-" gorm:"constraint:OnDelete:SET NULL;comment:删除模型的时候将引用模型的地方设置为null"`
	Dict            Dict     `json:"-"  gorm:"polymorphic:Dict;polymorphicValue:sys_go_type"`
	entity.Global
}
