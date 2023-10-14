package parameters

import "todo_list/internal/models"

type RegisterByEmailReq struct {
	Account      string `json:"account" binding:"required"`
	Password     string `json:"password" binding:"required"`
	AccountType  string `json:"accountType" gorm:"account_type" binding:"required"`
	Username     string `json:"username" binding:"required"`
	IdentifyCode string `json:"identifyCode" gorm:"-" binding:"required"`
	RegisterIp   string `json:"-" gorm:"colum:register_ip"`
}

type RegisterReq struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password"`
	AccountType  int    `json:"account_type" gorm:"account_type" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	IdentifyCode string `json:"identify_code" gorm:"-"`
	RegisterIp   string `json:"-" gorm:"colum:register_ip"`
	CreatedAt    int64  `json:"-" gorm:"autoCreateTime"`
	DeviceType   int    `json:"device_type" binding:"required"`
}
type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	LoginIp  string `json:"-"`
}
type NewLoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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

type Request interface {
	RegisterByEmailReq |
		CreateReq | LoginReq |
		InsertSuggestionReq | DeleteReq |
		ModifyUsernameReq | ResetPasswordReq |
		SendOTPReq | UpdateTaskReq |
		NewUpdateTaskReq | RegisterReq | NewLoginReq |
		models.TaskMode | models.TaskGroupModel
}
