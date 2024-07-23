package model

import (
	"xwya/model/global"
)

/**
+  code 类型表示
*  1 -------- 字典类型
*  2 -------- 框架选择
*  3 -------- GO语言类型
*  4 -------- Ts类型
*  5 -------- 求情方法
*/

type Dictionaries struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement;comment:自增id"`
	Code        uint       `json:"code" gorm:"type:varchar(50);not null;index:idx_code;comment:字典类型"` // 字典代码，唯一索引
	Value       string     `json:"value" gorm:"type:varchar(100);not null;index;comment:字典值"`         // 字典值
	Description string     `json:"description" gorm:"type:varchar(255);default:null;comment:描述"`      // 描述信息
	ProjectID   uint       `json:"project_id" gorm:"type:int;default:0;not null;comment:项目id"`
	Projects    *[]Project `json:"project,omitempty" gorm:"foreignKey:FrameID;constraint:OnDelete:SET	NULL" `
	global.Global
}
