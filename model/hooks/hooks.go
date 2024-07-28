package hooks

import (
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/sony/sonyflake"
)

var PassPrefix = "xwya"
var Flake *sonyflake.Sonyflake

func init() {
	fmt.Println("初始化雪花id生成器")
	Flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenerateID() uint {
	id, err := Flake.NextID()
	if err != nil {
		id = uint64(time.Now().Unix())
		// WriteLog(err)
		fmt.Println("生成雪花id失败", err)
	}
	return uint(id)
}
func EncryptPassword(password string) string {
	return cryptor.HmacMd5(PassPrefix, password)
}
