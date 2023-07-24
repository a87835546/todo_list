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
func (tl *UserLogic) QueryByAccount(req *parameters.LoginReq) (user *models.User, err error) {
	err = Db.Table("user").Where("account", req.Account).Find(&user).Error
	return
}
func (tl *UserLogic) UpdateIpByAccount(req *parameters.LoginReq) (err error) {
	err = Db.Table("user").Where("account=?", req.Account).Update("login_ip", req.LoginIp).Error
	return
}
