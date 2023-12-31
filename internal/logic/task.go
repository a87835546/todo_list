package logic

import (
	"encoding/json"
	"todo_list/internal/models"
	"todo_list/internal/parameters"
)

type TaskLogic struct {
}

func (tl *TaskLogic) Create(req *parameters.CreateReq) (task *models.Task, err error) {
	//one, err := MongoDB.Collection("task").InsertOne(context.Background(), req)
	//if err != nil {
	//	log.Println("err-->>", err.Error())
	//	return nil, err
	//}
	//log.Println("one --->>>>", one)
	err = Db.Table("task").Create(req).Error
	if err == nil {
		err = Db.Table("task").Where("name=?", req.TaskName).Find(&task).Error
	}
	return
}
func (tl *TaskLogic) CreateNew(req *models.TaskModel) (task *models.TaskModel, err error) {
	//one, err := MongoDB.Collection("task").InsertOne(context.Background(), req)
	//if err != nil {
	//	log.Println("err-->>", err.Error())
	//	return nil, err
	//}
	//log.Println("one --->>>>", one)
	err = Db.Table("task").Create(req).Error
	if err == nil {
		err = Db.Table("task").Where("id=?", req.Id).Find(&task).Error
	}
	return
}
func (tl *TaskLogic) QueryByUserId(id any) (task []*models.Task, err error) {
	err = Db.Table("task").Where("user_id=?", id).Find(&task).Error
	for i := 0; i < len(task); i++ {
		t := task[i]
		if len(t.TaskIcon) > 2 {
			icon := models.TaskIconModel{}
			err = json.Unmarshal([]byte(t.TaskIcon), &icon)
			t.TaskIconModel = icon
		}
		if len(t.Detail) > 2 {
			var list = make([]map[string]any, 0)
			err = json.Unmarshal([]byte(t.Detail), &list)
			var temp = make([]*models.TaskDetailModel, 0, len(list))
			for j := 0; j < len(list); j++ {
				tp := models.TaskDetailModel{}
				jsonData, _ := json.Marshal(list[j])
				err = json.Unmarshal(jsonData, &tp)
				temp = append(temp, &tp)
			}
			t.DetailModel = temp
		}
	}
	return
}
func (tl *TaskLogic) Delete(req *parameters.DeleteReq) (err error) {
	err = Db.Table("task").Delete("user_id=?", req.Account).Error
	return
}
func (tl *TaskLogic) Update(req *parameters.UpdateTaskReq) (err error) {
	err = Db.Table("task").Where("id=?", req.Id).Updates(req).Error
	return
}
func (tl *TaskLogic) NewUpdate(req *parameters.NewUpdateTaskReq) (err error) {
	err = Db.Table("task").Where("id=?", req.Id).Updates(req).Error
	return
}
func (tl *TaskLogic) QueryTasksNew(userId any) (list []*models.TaskModel, err error) {
	err = Db.Table("task").Where("user_id=?", userId).Find(&list).Error
	return
}
func (tl *TaskLogic) QueryTasksCountNew(userId any) (list []*models.TaskCountModel, err error) {
	err = Db.Table("task").Debug().Raw("select tmp.*,task_group.name,task_group.name_en,task_group.color,task_group.icon from(select count(task_group_id) as count,task_group_id from task group by task_group_id HAVING task_group_id in(select id from task_group where user_id in(?,0)))as tmp left join task_group on tmp.task_group_id=task_group.id", userId).Find(&list).Error
	var temp []*models.TaskModel
	err = Db.Table("task").Debug().Raw("select *\nfrom task\nwhere task_group_id in (select id from task_group where user_id in (?,0))", userId).Find(&temp).Error
	if err == nil {
		for i := 0; i < len(list); i++ {
			m := list[i]
			for j := 0; j < len(temp); j++ {
				m1 := temp[j]
				if m.Id == m1.TaskGroupId {
					m.Tasks = append(m.Tasks, m1)
				}
			}
		}
	}
	return
}
