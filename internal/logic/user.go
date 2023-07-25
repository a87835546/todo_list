package logic

import (
	"log"
	"todo_list/internal/models"
	"todo_list/internal/parameters"
)

type UserLogic struct {
}

func (tl *UserLogic) Create(req *parameters.RegisterByEmailReq) (user *models.User, err error) {
	err = Db.Table("user").Create(req).Error
	log.Printf("插入数据异常-->>> %#v", err)
	if err == nil {
		err = Db.Table("user").Where("account", req.Account).Find(&user).Error
	}
	return
}
func (tl *UserLogic) QueryByAccount(account string) (user *models.User, err error) {
	err = Db.Table("user").Where("account", account).Find(&user).Error
	return
}
func (tl *UserLogic) UpdateUsernameById(req *parameters.ModifyUsernameReq) (err error) {
	err = Db.Table("user").Where("account=?", req.Account).Update("username", req.UserName).Error
	return
}

func (tl *UserLogic) UpdatePasswordById(req *parameters.ResetPasswordReq) (err error) {
	err = Db.Table("user").Where("account=?", req.Account).Update("password", req.NewPassword).Error
	return
}

func (tl *UserLogic) UpdateIpByAccount(req *parameters.LoginReq) (err error) {
	err = Db.Table("user").Where("account=?", req.Account).Update("login_ip", req.LoginIp).Error
	return
}

func (tl *UserLogic) CreateSuggestion(req *parameters.InsertSuggestionReq) (err error) {
	err = Db.Table("suggestion").Create(req).Error
	log.Printf("插入数据异常-->>> %#v", err)
	return
}
func (tl *UserLogic) QuerySuggestionByAccount(id int) (user *models.Suggestion, err error) {
	err = Db.Table("suggestion").Where("user_id", id).Find(&user).Error
	return
}
