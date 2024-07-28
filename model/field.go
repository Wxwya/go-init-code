package model

import (
	"xwya/entity"
)

/*
		References 字段需要设置当前表中的字段
	  ForeignKey 设置模型表中的字段
*/
type Field struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement;comment:字段ID"`
	FormID           *uint  `json:"form_id" gorm:"default:0;index;comment:表单的id"`
	FieldTypeID      uint   `json:"field_type_id" gorm:"not null;comment:字段类型(通过字典sys_go_type获取)"`
	FieldName        string `json:"field_name" gorm:"type:varchar(30);not null;comment:字段名称"`
	FieldSize        *uint  `json:"field_size" gorm:"comment:字段大小"`
	IsNull           *bool  `json:"is_null" gorm:"default:0;comment:是否为null"`
	AutoIncrement    *bool  `json:"auto_increment" gorm:"default:0;comment:是否自增"`
	IsPrimaryKey     *bool  `json:"is_primary_key" gorm:";default:0;comment:是否为主键"`
	IsUnique         *bool  `json:"is_unique" gorm:"not null;default:0;comment:是否唯一值"`
	ForeignKey       *bool  `json:"foreign_key" gorm:"not null;default:0;comment:设置外键"`
	References       string `json:"references" gorm:"comment:设置映射"`
	IndexName        string `json:"index_name" gorm:"comment:索引名称(为复合索引使用)"`
	IsValNull        *bool  `json:"is_val_null" gorm:"default:true;comment:为空时该字段是否返回"`
	GlobalModelID    *uint  `json:"global_model_id" gorm:"index;defautl:0;comment:全局模型ID"`
	IsQuery          *bool  `json:"is_query" gorm:"default:false;comment:是否是查询字段"`
	IsPointer        *bool  `json:"is_pointer" gorm:"default:false;comment:是否是指针"`
	IsAdd            *bool  `json:"is_add" gorm:"default:false;comment:是否是添加字段"`
	JoinForeignKey   string `json:"join_foreign_key" gorm:"comment:外键字段"`
	JoinReferences   string `json:"join_references" gorm:"comment:映射字段"`
	Many             string `json:"many" gorm:"comment:关联模型(用于多对多)"`
	Polymorphic      string `json:"polymorphic" gorm:"comment:多态关联模型" `
	PolymorphicValue string `json:"polymorphic_value" gorm:"comment:多态关联模型值"`
	Constraint       *bool  `json:"constraint" gorm:"default:false;comment:是否是约束字段"`
	DefaultValue     string `json:"default_value" gorm:"comment:默认值"`
	OnUpdate         string `json:"on_update" gorm:"comment:更新时触发(字典sys_con_type)"`
	OnDelete         string `json:"on_delete" gorm:"comment:删除时触发(字典sys_con_type)"`
	Check            string `json:"check" gorm:"comment:设置字段约束"`
	Embedded         *bool  `json:"embedded" gorm:"default:false;comment:是否是嵌入字段"`
	EmbeddedPrefix   string `json:"embedded_prefix" gorm:"comment:嵌入字段前缀"`
	Precision        string `json:"precision" gorm:"comment:设置数字精度"`
	Scale            string `json:"scale" gorm:"comment:设置数字小数位数"`
	entity.Global
}

type FieldDto struct {
	Field
	FieldTypeName string `json:"field_type_name"`
}
