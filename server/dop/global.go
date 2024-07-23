package dop

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"gorm.io/gorm"
	"xwya/utils"
)

func AddPagination(db *gorm.DB, pageNum, pageSize int, order string) {
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("create_time " + order)
}

// 加密密码
func EncryptPassword(password string) string {
	return cryptor.HmacMd5(utils.PassPrefix, password)
}
