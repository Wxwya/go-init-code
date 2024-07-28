package sqltype

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime time.Time

func (t *CustomTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *CustomTime) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*t = CustomTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
