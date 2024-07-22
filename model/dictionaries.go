package model

import (
	"xwya/model/global"

	"gorm.io/gorm"
)

/**
+  code 类型表示
*  1 -------- 字典类型
*  2 -------- 框架选择
*  3 -------- GO语言类型
*  4 -------- Ts类型
*/

type Dictionaries struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:自增id"`
	Code        uint   `json:"code" gorm:"type:varchar(50);not null;index:idx_code;comment:字典类型"` // 字典代码，唯一索引
	Value       string `json:"value" gorm:"type:varchar(100);not null;index;comment:字典值"`         // 字典值
	Description string `json:"description" gorm:"type:varchar(255);default:null;comment:描述"`      // 描述信息
	global.Global
}

type FuncType func(tx *gorm.DB, d *Dictionaries) error

var handleUpdate = map[uint]FuncType{
	2: updateProjectFrameName,
}

func updateProjectFrameName(tx *gorm.DB, d *Dictionaries) error {

	if existingValue, ok := tx.Statement.Get("existingValue"); ok && existingValue != d.Value {
		if err := tx.Model(&Project{}).Where(&Project{FrameID: d.ID}).Update("frame_name", d.Value).Error; err != nil {
			return err
		}
	}
	return nil
}
func (d *Dictionaries) BeforeUpdate(tx *gorm.DB) error {
	var existingValue string
	if err := tx.Model(&Dictionaries{}).Select("Value").Where(&Dictionaries{ID: d.ID}).Scan(&existingValue).Error; err != nil {
		return err
	}
	tx.Statement.Set("existingValue", existingValue)
	return nil
}

func (d *Dictionaries) AfterUpdate(tx *gorm.DB) error {
	var f, ok = handleUpdate[d.Code]
	if !ok {
		return nil
	}
	if err := f(tx, d); err != nil {
		return err
	}

	return nil
}
