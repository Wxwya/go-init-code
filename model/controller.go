package model

import "xwya/model/global"

/*
	控制器(对应api文件夹)
*/
type Controller struct {
	ID          int                    `json:"id" gorm:"primaryKey;autoIncrement;comment:表ID"`
	Version     string                 `json:"version" gorm:"type:varchar(10);not null;comment:控制器版本号"`
	Name        string                 `json:"name" gorm:"type:varchar(30);not null;unique;comment:文件名称"`
	Description string                 `json:"description" gorm:"type:text;comment:文件描述"`
	Processors  *[]ControllerProcessor `json:"processors,omitempty" gorm:"foreignKey:ControllerID;constraint:OnDelete:CASCADE;comment:一对多关系"`
	global.Global
}
