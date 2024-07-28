package sqltype

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/slice"
)

type StringArray []string

func (a *StringArray) Scan(val any) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), ",")
	if len(ss) == 1 && ss[0] == "" {
		*a = []string{}
	} else {
		*a = ss
	}
	return nil
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type IntArray []int

func (a *IntArray) Scan(val any) error {
	str := string(val.([]uint8))
	arr := strings.Split(str, ",")
	if len(arr) == 1 && arr[0] == "" {
		*a = []int{}
	} else {
		*a = slice.FlatMap(arr, func(i int, num string) []int {
			aa, _ := strconv.Atoi(num)
			return []int{aa}
		})
	}
	// json.Unmarshal(val.([]uint8), a)
	return nil
}

// Value 实现 driver.Valuer 接口
func (a IntArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}
