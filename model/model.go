package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id" ` // 主键ID
	CreatedAt LocalTime      `json:"created_at"`                           // 创建时间
	UpdatedAt LocalTime      `json:"-"`                                    // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                       // 删除时间
}

const (
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)

type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// ""空值不进行解析
	if len(data) == 2 {
		*t = LocalTime{Time: time.Time{}}
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+SecLocalTimeFormat+`"`, string(data))
	*t = LocalTime{Time: now}
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(SecLocalTimeFormat))
	return []byte(output), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	/*
		gorm 写入 mysql 时调用
	*/
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	/*
		gorm 检出 mysql 时调用
	*/
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to LocalTime", v)
}

func (t LocalTime) String() string {
	/*
		用于 fmt.Println 和后续验证场景
	*/
	return t.Format(SecLocalTimeFormat)
}

func (t LocalTime) DateString() string {
	/*
		只需要日期
	*/
	return t.Format(DateLocalTimeFormat)
}
