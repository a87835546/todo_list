package models

type Time struct {
	Id        int64 `json:"id,omitempty"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}
type UserModel struct {
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"-"`
	Avatar      string `json:"avatar,omitempty"`
	LoginIp     string `json:"-" gorm:"colum:login_ip"`
	RegisterIp  string `json:"-" gorm:"colum:register_ip"`
	IsDelete    int    `json:"-" gorm:"colum:is_delete"`
	AccountType int    `json:"account_type" gorm:"colum:account_type"`
	DeviceType  int    `json:"device_type" gorm:"colum:device_type"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Time
}

type TaskGroupModel struct {
	Color  int64  `json:"color"`
	Icon   string `json:"icon"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	UserId int    `json:"user_id"`
	Time
}

type TaskModel struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Type        int    `json:"type"`
	TaskGroupId int64  `json:"task_group_id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Process     int    `json:"process"`
	EndTime     int64  `json:"end_time"`
	StartTime   int64  `json:"start_time"`
	ProjectName string `json:"project_name"`
	Time
}

type TaskCountModel struct {
	Count  int    `json:"count"`
	Id     int    `json:"task_group_id" gorm:"column:task_group_id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Color  int64  `json:"color"`
	Icon   string `json:"icon"`
}
