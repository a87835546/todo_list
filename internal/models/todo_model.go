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
	LoginIp     string `json:"-"`
	RegisterIp  string `json:"-"`
	IsDelete    int    `json:"-"`
	AccountType int    `json:"account_type,omitempty"`
	DeviceType  int    `json:"device_type,omitempty"`
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

type TaskMode struct {
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
