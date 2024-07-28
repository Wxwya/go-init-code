package model

import "xwya/entity"

/*
	对应msg文件夹
	创建表时需要生成自动生成 success,error,serverError,以及token的错误信息,用户登录的信息
*/
type MsgCode struct {
	ID      uint                 `json:"id" gorm:"primaryKey;autoIncrement;comment:表ID"`
	MsgKey  string               `json:"msg_key" gorm:"not null;size:50;comment:msg的key"`
	Code    uint                 `json:"code" gorm:"not null;comment:msg的状态码"`
	Msg     string               `json:"msg" gorm:"not null;comment:返回的内容"`
	Success *ControllerProcessor `json:"-" gorm:"foreignKey:SuccessCodeID;references:ID; constraint:OnDelete:SET NULL;comment:一对多关系"`
	Error   *ControllerProcessor `json:"-" gorm:"foreignKey:SuccessCodeID;references:ID; constraint:OnDelete:SET NULL;comment:一对多关系"`
	entity.Global
}
