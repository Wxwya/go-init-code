package model

import (
	"xwya/entity"
)

/**
+  DiceType 类型表示
*/

type Dict struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:自增id"`
	Code        string `json:"label" gorm:"type:varchar(50);not null;index:idx_code;comment:字典的键(label)"` // 字典代码，唯一索引
	Description string `json:"description" gorm:"type:varchar(255);default:null;comment:描述"`              // 描述信息
	ProjectID   uint   `json:"project_id" gorm:"not null;index;default:0;comment:项目id(为0的话代表系统默认的值)"`
	DictID      uint   `json:"dict_id" gorm:"not null;index;comment:字典所属的表的id(字典值value)"`
	DictType    string `json:"dict_type" gorm:"not null;index;comment:字典类型(查询使用)"`
	entity.Global
}
