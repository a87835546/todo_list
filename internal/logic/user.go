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
func (tl *UserLogic) QuerySuggestionByAccount(id int) (list []*models.Suggestion, err error) {
	err = Db.Table("suggestion").Where("user_id", id).Find(&list).Error
	return
}

func (tl *UserLogic) CreateAccount(req *parameters.RegisterReq) (user *models.UserModel, err error) {
	user, err = tl.QueryUserByEmail(req.Email)
	if err == nil && user.Id > 0 {
		return user, nil
	} else {
		err = Db.Debug().Table("user").Create(req).Error
		log.Printf("插入数据异常-->>> %#v", err)
		if err == nil {
			user, err = tl.QueryUserByEmail(req.Email)
		}
	}
	return
}
func (tl *UserLogic) QueryUserByEmail(email string) (user *models.UserModel, err error) {
	err = Db.Debug().Table("user").Where("email=?", email).Find(&user).Error
	return
}
func (tl *UserLogic) Update(user *models.UserModel) (err error) {
	err = Db.Debug().Table("user").Updates(user).Error
	return
}
