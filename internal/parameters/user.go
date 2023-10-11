package parameters

type RegisterByEmailReq struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	AccountType  string `json:"accountType" gorm:"account_type"`
	Username     string `json:"username"`
	IdentifyCode string `json:"identifyCode" gorm:"-"`
	RegisterIp   string `json:"-" gorm:"colum:register_ip"`
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
