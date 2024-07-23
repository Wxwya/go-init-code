package global

import (
	"xwya/utils/sqltype"
)

type Global struct {
	CreateTime sqltype.CustomTime `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	UpdateTime sqltype.CustomTime `json:"-" gorm:"type:datetime;autoUpdateTime"`
}
type Page struct {
	PageSize int `json:"pagesize" binding:"required"`
	PageNum  int `json:"pagenum"  binding:"required,min=1"`
}
