package model

import (
	"xwya/model/global"

	"gorm.io/gorm"
)

type Form struct {
	ID          uint     `json:"id" gorm:"primaryKey;autoIncrement;comment:表ID"`
	ProjectID   uint     `json:"project_id"  gorm:"index:idx_form_projeceid_formname;comment:项目ID"`
	FormName    string   `json:"form_name" gorm:"index:idx_form_projeceid_formname;size:100;not null;comment:表名"`
	FormComment string   `json:"form_comment" gorm:"comment:描述"`
	IsChange    *bool    `json:"is_change,omitempty"  gorm:"default:false;comment:判断是否修改"`
	Fields      *[]Field `json:"fields,omitempty" gorm:"constraint:OnDelete:CASCADE;comment:一对多关系"`
	global.Global
}

func (f *Form) AfterFind(tx *gorm.DB) (err error) {
	f.IsChange = nil
	return
}
