package models

import "time"

type Log struct {
	Id        int64     `json:"id"`
	Type      int       `json:"type"` // 日志类型 1---登录
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	Desc      string    `json:"desc"`
}
