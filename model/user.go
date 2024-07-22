package model

import (
	"xwya/model/global"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(30);not null"`
	Account  string `json:"account" gorm:"type:varchar(50);not null"`
	Password string `json:"password,omitempty" gorm:"type:varchar(60);not null"`
	// Account
	global.Global
}
