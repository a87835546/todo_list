package logic

import (
	"todo_list/internal/models"
	"todo_list/internal/parameters"
)

type TaskLogic struct {
}

func (tl *TaskLogic) Create(req *parameters.CreateReq) (task *models.Task, err error) {
	err = Db.Table("task").Create(req).Error
	if err == nil {
		err = Db.Table("task").Where("taskName=?", req.TaskName).Find(&task).Error
	}
	return
}
