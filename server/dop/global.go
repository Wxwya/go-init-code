package dop

import (
	"fmt"
	"reflect"
	"strconv"
	"xwya/utils"

	"github.com/duke-git/lancet/v2/cryptor"
	"gorm.io/gorm"
)

func IsBool(db *gorm.DB, key string, flag *bool) {
	whereStr := fmt.Sprintf("%s = ?", key)
	if flag != nil {
		db.Where(whereStr, *flag)
	}
}
func IsInt(db *gorm.DB, key string, v interface{}) {
	var i float64
	whereStr := fmt.Sprintf("%s = ?", key)

	switch v := v.(type) {
	case uint:
		i = float64(v)
	case int:
		i = float64(v)
	case float64:
		i = v
	case string:
		i, _ = strconv.ParseFloat(v, 64)
	default:
		// 处理不支持的类型
		fmt.Printf("Value: %v, Type: %s\n", v, reflect.TypeOf(v))
		panic("unsupported type for v")
	}
	if i != 0 {
		db.Where(whereStr, i)
	}
}

func AddPagination(db *gorm.DB, pageNum, pageSize int, order string) {
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("create_time " + order)
}

// 加密密码
func EncryptPassword(password string) string {
	return cryptor.HmacMd5(utils.PassPrefix, password)
}
