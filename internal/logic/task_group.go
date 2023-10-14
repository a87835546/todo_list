package logic

import "todo_list/internal/models"

type TaskGroupService struct {
}

func NewTaskGroupService() *TaskGroupService {
	return &TaskGroupService{}
}
func (TaskGroupService) QueryById(id int64) (model models.TaskGroupModel, err error) {
	err = Db.Table("task_group").Debug().Where("id=?", id).Find(&model).Error
	return
}

func (TaskGroupService) QueryByUserId(id int) (model []*models.TaskGroupModel, err error) {
	err = Db.Table("task_group").Debug().Where("user_id=? ", []int{id, 0}).Find(&model).Error
	return
}
func (TaskGroupService) Add(model *models.TaskGroupModel) (err error) {
	err = Db.Table("task_group").Debug().Create(&model).Error
	return
}
func (TaskGroupService) Update(model *models.TaskGroupModel) (err error) {
	err = Db.Table("task_group").Debug().Updates(&model).Error
	return
}
