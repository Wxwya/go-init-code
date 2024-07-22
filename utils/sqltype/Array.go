package sqltype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (a *Array) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), ",")
	if len(ss) == 1 && ss[0] == "" {
		*a = []string{}
	} else {
		*a = ss
	}
	return nil
}

func (a Array) Value() (driver.Value, error) {
	str := strings.Join(a, ",")
	return str, nil
}
