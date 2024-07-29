package model

import (
	"测试项目1/enetity"
	"测试项目1/enetity/sqltype"
)

type BusinessUser struct {
	Status        *bool              `json:"status" 	gorm:"not null;index;" `
	Id            string             `json:"id" 	gorm:"primaryKey;autoIncrement;" `
	Username      string             `json:"username" 	gorm:"index;size:60;" `
	Account       string             `json:"account" 	gorm:"index;size:60;" `
	Password      string             `json:"password" 	gorm:"size:60;" `
	Avatar        string             `json:"avatar" 	gorm:"" `
	Email         string             `json:"email" 	gorm:"size:30;" `
	Phone         string             `json:"phone" 	gorm:"size:20;" `
	LastLoginTime sqltype.CustomTime `json:"last_login_time" 	gorm:"" `
	Role          []Role             `json:"-" 	gorm:"index;many2many:business_user_role;" `
	enetity.Global
}
