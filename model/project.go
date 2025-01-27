package model

import (
	"xwya/entity"
)

type Project struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectName string `json:"project_name" gorm:"index:idx_projectname_sqlhost_frame_id;not null;comment:项目名称"`
	Description string `json:"description" gorm:"not null;comment:描述"`
	SqlHost     string `json:"sql_host" gorm:"index:idx_projectname_sqlhost_frame_id;type:varchar(30);not null;comment:数据库地址"`
	SqlProt     string `json:"sql_prot" gorm:"type:varchar(10);not null;comment:数据库端口"`
	SqlBaseName string `json:"sql_basename" gorm:"type:varchar(50);not null;comment:数据库库名"`
	SqlUser     string `json:"sql_user" gorm:"type:varchar(30);not null;comment:数据库账号"`
	SqlPassword string `json:"sql_password" gorm:"type:varchar(30);not null;comment:数据库密码"`
	FrameType   string `json:"frame_type,omitempty" gorm:"index:idx_projectname_sqlhost_frame_id;comment:前端使用的端架"`
	Forms       []Form `json:"forms,omitempty" gorm:"constraint:OnDelete:CASCADE;comment:一对多关系"`
	entity.Global
}
