package model

import (
	"xwya/entity"

	"gorm.io/gorm"
)

/*
*
建表时需要自动创建类型到字典表 code 3 中
*/
type Form struct {
	ID            uint    `json:"id" gorm:"primaryKey;autoIncrement;comment:表ID"`
	ProjectID     uint    `json:"project_id"  gorm:"index:idx_form_projeceid_formname;comment:项目ID"`
	FormName      string  `json:"form_name" gorm:"index:idx_form_projeceid_formname;size:100;not null;comment:表名"`
	FormComment   string  `json:"form_comment" gorm:"comment:描述"`
	IsChange      *bool   `json:"is_change,omitempty"  gorm:"default:false;comment:判断是否修改"`
	Fields        []Field `json:"fields,omitempty" gorm:"constraint:OnDelete:CASCADE;comment:一对多关系"`
	BeforeCreates string  `json:"before_creates" gorm:"type:text;comment:钩子函数创建前"`
	AfterCreates  string  `json:"after_creates" gorm:"type:text;comment:钩子函数创建后"`
	BeforeSaves   string  `json:"before_saves" gorm:"type:text;comment:钩子函数保存前"`
	AfterSaves    string  `json:"after_saves" gorm:"type:text;comment:钩子函数保存后"`
	BeforeUpdates string  `json:"before_updates" gorm:"type:text;comment:钩子函数更新前"`
	AfterUpdates  string  `json:"after_updates" gorm:"type:text;comment:钩子函数更新后"`
	BeforeDeletes string  `json:"before_deletes" gorm:"type:text;comment:钩子函数删除前"`
	AfterDeletes  string  `json:"after_deletes" gorm:"type:text;comment:钩子函数删除后"`
	AfterFinds    string  `json:"after_finds" gorm:"type:text;comment:钩子函数查询后"`
	Dict          []Dict  `json:"-" gorm:"polymorphic:Dict;polymorphicValue:sys_go_type"`
	entity.Global
}

func (f *Form) AfterFind(tx *gorm.DB) (err error) {
	f.IsChange = nil
	return
}
