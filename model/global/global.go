package global

import (
	f "xwya/utils/sqltype"
)

type Global struct {
	CreateTime f.CustomTime `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	UpdateTime f.CustomTime `json:"-" gorm:"type:datetime;autoUpdateTime"`
}
type Page struct {
	PageSize int `json:"pagesize"`
	PageNum  int `json:"pagenum"`
}
