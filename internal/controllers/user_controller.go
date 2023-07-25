package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/parameters"
)

type UserCtl struct {
}

var ul = logic.UserLogic{}

func (us *UserCtl) Login(ctx *gin.Context) {
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

func (us *UserCtl) UpdateUserName(ctx *gin.Context) {
	req := parameters.ModifyUsernameReq{}
	ParserReqParameters(&req, ctx)
	err := ul.UpdateUsernameById(&req)
	if err != nil {
		RespError(ctx, UpdateDBErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}

func (us *UserCtl) UploadAvatar(ctx *gin.Context) {
	req := parameters.ModifyUsernameReq{}
	ParserReqParameters(&req, ctx)
	err := ul.UpdateUsernameById(&req)
	if err != nil {
		RespError(ctx, UpdateDBErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}

func (us *UserCtl) IdentifyCodeSend(ctx *gin.Context) {
	req := parameters.SendOTPReq{}
	ParserReqParameters(&req, ctx)
	otp := GenValidateCode(6)
	key := "otp:account:" + req.Account
	err := logic.Client.Set(key, otp, 5*60*1000).Err()
	if err != nil {
		RespError(ctx, SendOtpErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}

func (us *UserCtl) Register(ctx *gin.Context) {
	req := parameters.RegisterByEmailReq{}
	ParserReqParameters(&req, ctx)
	log.Printf("req--->>> %s", req)
	req.RegisterIp = GetRequestIP(ctx)
	user, err := ul.Create(&req)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, err.Error())
	} else {
		generateToken(ctx, user)
	}
}

func (us *UserCtl) ResetPassword(ctx *gin.Context) {
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

func (us *UserCtl) ForgetPassword(ctx *gin.Context) {
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

func (us *UserCtl) OneDaySuggestion(ctx *gin.Context) {
	req := parameters.InsertSuggestionReq{}
	ParserReqParameters(&req, ctx)
	err := ul.CreateSuggestion(&req)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}

func (us *UserCtl) GetSuggestion(ctx *gin.Context) {
	list, err := ul.QuerySuggestionByAccount(1)
	if err != nil {
		RespError(ctx, QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, list)
	}
}
