package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/models"
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
	user, err := ul.QueryByAccount(&req)
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
	RespOk(ctx, models.User{})
}

func (us *UserCtl) UploadAvatar(ctx *gin.Context) {
	RespOk(ctx, models.User{})
}

func (us *UserCtl) IdentifyCodeSend(ctx *gin.Context) {
	RespOk(ctx, models.User{})
}

func (us *UserCtl) IdentifyCodeCheck(ctx *gin.Context) {
	RespOk(ctx, models.User{})
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
	RespOk(ctx, models.User{})
}

func (us *UserCtl) ForgetPassword(ctx *gin.Context) {
	RespOk(ctx, models.User{})
}

func (us *UserCtl) OneDaySuggestion(ctx *gin.Context) {
	RespOk(ctx, models.User{})
}

func (us *UserCtl) GetSuggestion(ctx *gin.Context) {
	RespOk(ctx, []models.User{})
}
