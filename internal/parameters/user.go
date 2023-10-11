package parameters

type Time struct {
	Id        int64 `json:"id,omitempty"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}
type RegisterByEmailReq struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	AccountType  string `json:"accountType" gorm:"account_type"`
	Username     string `json:"username"`
	IdentifyCode string `json:"identifyCode" gorm:"-"`
	RegisterIp   string `json:"-" gorm:"colum:register_ip"`
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

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	LoginIp  string `json:"-"`
}

type InsertSuggestionReq struct {
	AvatarUrl  string `json:"avatar" gorm:"-"`
	Account    string `json:"account" gorm:"column:user_id"`
	Suggestion string `json:"suggestion" gorm:"column:title"`
	ConnectWay string `json:"connectWay" gorm:"column:desc"`
	UserName   string `json:"userName" gorm:"column:username"`
}

type ModifyUsernameReq struct {
	Account  string `json:"account"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}

type ResetPasswordReq struct {
	Account         string `json:"account"`
	Token           string `json:"token"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ForgetPasswordReq struct {
	Account         string `json:"account"`
	AccountType     string `json:"accountType"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
	IdentifyCode    string `json:"identifyCode"`
}

type SendOTPReq struct {
	Account  string
	Why      string
	Language string
	//"account": widget.account??"",
	//"why": widget.isForgetPassword??false ? "emailForget" : "emailRegister",
	//"language": globalModel.currentLanguageCode[0]
}
