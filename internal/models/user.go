package models

import "time"

type User struct {
	Id          int       `json:"id"`
	AccountType int       `json:"account_type" gorm:"account_type"`
	Username    string    `json:"username"`
	Account     string    `json:"account"` // 账号(email或者phone)
	Avatar      string    `json:"avatar"`
	Position    string    `json:"position"`
	Password    string    `json:"password"`
	DeviceType  int       `json:"device_type" gorm:"device_type"` // 0--unknown/1--ios/2--android
	LoginIp     string    `json:"login_ip" gorm:"login_ip"`
	RegisterIp  string    `json:"register_ip" gorm:"register_ip"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}
