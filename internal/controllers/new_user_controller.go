package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/parameters"
)

type NewUserController struct {
}

func New() *NewUserController {
	return &NewUserController{}
}

func (us *NewUserController) Query(ctx *gin.Context) {
	RespOk(ctx, nil)
}
func (us *NewUserController) Login(ctx *gin.Context) {
	req := parameters.LoginReq{}
	ParserReqParameters(&req, ctx)
	log.Printf("req--->>> %s", req)
	req.LoginIp = GetRequestIP(ctx)
	user, err := ul.QueryByAccount(req.Account)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, err.Error())
	} else if user.Password != req.Password {
		RespError(ctx, LoginPasswordErrorCode, "密码或账号错误")
	} else {
		err = ul.UpdateIpByAccount(&req)
		if err != nil {
			log.Println("更新登录ip异常", err.Error())
		}
		generateToken(ctx, user)
	}
}
func (us *NewUserController) Register(ctx *gin.Context) {
	req := parameters.RegisterReq{}
	err := ParserReq(&req, ctx)
	if err != nil {
		RespErrorWithMsg(ctx, ParameterErrorCode, "解析参数异常"+err.Error(), nil)
	} else {
		log.Printf("req--->>> %s", req)
		req.RegisterIp = GetRequestIP(ctx)
		user, err := ul.CreateAccount(&req)
		if err != nil {
			RespError(ctx, InsertDBErrorCode, err.Error())
		} else {
			generateNewToken(ctx, user)
		}
	}
}

func (us *NewUserController) NewLogin(ctx *gin.Context) {
	req := parameters.NewLoginReq{}
	err := ParserReq(&req, ctx)
	if err != nil {
		RespErrorWithMsg(ctx, ParameterErrorCode, "解析参数异常"+err.Error(), nil)
	} else {
		log.Printf("req--->>> %s", req)
		req.LoginIp = GetRequestIP(ctx)
		user, err := ul.QueryUserByEmail(req.Email)
		if err != nil {
			RespError(ctx, InsertDBErrorCode, err.Error())
		} else {
			if user.Password == req.Password {
				user.LoginIp = GetRequestIP(ctx)
				err = ul.Update(user)
				if err != nil {
					RespError(ctx, InsertDBErrorCode, err.Error())
				}
			}
			generateNewToken(ctx, user)
		}
	}
}

func (us *NewUserController) ResetPassword(ctx *gin.Context) {
	req := parameters.ResetPasswordReq{}
	ParserReqParameters(&req, ctx)
	user, err := ul.QueryByAccount(req.Account)
	if user.Password != req.OldPassword {
		RespError(ctx, UpdatePasswordErrorCode, "旧密码不匹配")
	} else {
		err = ul.UpdatePasswordById(&req)
		if err != nil {
			RespError(ctx, UpdateDBErrorCode, err.Error())
		} else {
			RespOk(ctx, nil)
		}
	}
}

func (us *NewUserController) ForgetPassword(ctx *gin.Context) {
	req := parameters.ModifyUsernameReq{}
	ParserReqParameters(&req, ctx)
	key := "otp:account:" + req.Account
	res, err := logic.Client.Get(key).Result()
	if res != req.Token {
		RespError(ctx, OtpErrorCode, "验证码异常")
	} else {
		err = ul.UpdateUsernameById(&req)
		if err != nil {
			RespError(ctx, UpdateDBErrorCode, err.Error())
		} else {
			RespOk(ctx, nil)
		}
	}
}
func (us *NewUserController) GetOTPByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	key := "otp:account:" + email
	res, err := logic.Client.Set(key, "", 5*60).Result()
	if err != nil {
		RespError(ctx, OtpErrorCode, err.Error())
	} else {
		RespOk(ctx, res)
	}
}
