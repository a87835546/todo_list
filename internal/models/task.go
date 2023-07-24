package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id              int64              `json:"id"`
	UserId          int                `json:"user_id" gorm:"user_id"`
	UserAccount     string             `json:"user_account" gorm:"user_account"`
	CreatedAt       time.Time          `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" gorm:"updated_at"` // on update now()
	Name            string             `json:"name"`
	Type            int                `json:"type"` // 任务类型。
	Status          int                `json:"status"`
	StartDate       time.Time          `json:"start_date" gorm:"start_date"`
	EndDate         time.Time          `json:"end_date" gorm:"end_date"`
	DeadLineDate    time.Time          `json:"dead_line_date" gorm:"dead_line_date"`
	OverallProgress int                `json:"overall_progress" gorm:"overall_progress"`
	TaskIcon        string             `json:"-" gorm:"task_icon" gorm:"type:json" gorm:"embedded"`
	TaskIconModel   TaskIconModel      `json:"task_icon" gorm:"-"`
	Detail          string             `json:"-" gorm:"column:detail"`
	DetailModel     []*TaskDetailModel `json:"detail" gorm:"-"`
	DetailNum       int                `json:"detail_num" gorm:"detail_num"`
}

type TaskDetailModel struct {
	TaskDetailName string `json:"taskDetailName"`
	ItemProgress   int    `json:"itemProgress"`
}
type TaskIconModel struct {
	TaskName string      `json:"taskName"`
	Icon     *IconModel  `json:"iconBean" gorm:"embedded"`
	Color    *ColorModel `json:"colorBean" gorm:"embedded"`
}

type IconModel struct {
	CodePoint          int    `json:"codePoint,omitempty"`
	FontFamily         string `json:"fontFamily"`
	FontPackage        string `json:"fontPackage"`
	IconName           string `json:"iconName"`
	MatchTextDirection string `json:"matchTextDirection"`
}

type ColorModel struct {
	Red     int `json:"red"`
	Green   int `json:"green"`
	Blue    int `json:"blue"`
	Opacity int `json:"opacity"`
}

type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
