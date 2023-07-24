package parameters

type RegisterByEmailReq struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	AccountType  string `json:"accountType" gorm:"account_type"`
	Username     string `json:"username"`
	IdentifyCode string `json:"identifyCode" gorm:"-"`
	RegisterIp   string `json:"-" gorm:"register_ip"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	LoginIp  string `json:"-"`
}
