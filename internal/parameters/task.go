package parameters

type CreateReq struct {
	TaskName        string `json:"taskName" gorm:"column:name" db:"name"`
	TaskType        int    `json:"taskType" gorm:"column:type"`
	Account         string `json:"account" gorm:"column:user_account"`
	UserId          int    `json:"user_id" gorm:"column:user_id"`
	TaskStatus      int    `json:"taskStatus" gorm:"column:status"`
	TaskDetailNum   int    `json:"taskDetailNum" gorm:"column:detail_num"`
	OverallProgress int    `json:"overallProgress" gorm:"column:overall_progress"`
	ChangeTimes     string `json:"changeTimes" gorm:"column:updated_at"`
	FinishDate      string `json:"finishDate" gorm:"column:end_date"`
	StartDate       string `json:"startDate" gorm:"column:start_date"`
	DeadLine        string `json:"deadLine" gorm:"column:dead_line_date"`
	TaskIconBean    string `json:"taskIconBean" gorm:"column:task_icon"`
	DetailList      string `json:"detailList" gorm:"column:detail"`
	Token           string `json:"token" gorm:"-"`
}

type DeleteReq struct {
	Token   string `json:"token"`
	Account string `json:"account"`
}

type UpdateTaskReq struct {
	TaskName        string `json:"taskName,omitempty" gorm:"column:name"`
	TaskType        int    `json:"taskType,omitempty" gorm:"column:type"`
	Account         string `json:"account,omitempty"`
	TaskStatus      int    `json:"taskStatus,omitempty"`
	TaskDetailNum   int    `json:"taskDetailNum,omitempty"`
	OverallProgress int    `json:"overallProgress,omitempty" gorm:"column:overall_progress"`
	ChangeTimes     string `json:"changeTimes,omitempty" gorm:"column:updated_at"`
	FinishDate      string `json:"finishDate,omitempty" gorm:"column:end_date"`
	StartDate       string `json:"startDate,omitempty" gorm:"column:start_date"`
	Id              string `json:"uniqueId,omitempty"`
	DeadLine        string `json:"deadLine,omitempty" gorm:"column:dead_line_date"`
	TaskIconBean    string `json:"taskIconBean,omitempty" gorm:"column:task_icon"`
	DetailList      string `json:"detailList,omitempty" gorm:"detail_list"`
	Token           string `json:"token,omitempty" gorm:"-"`

	//'taskName':taskBean.taskName,
	//'taskType':taskBean.taskType,
	//'account':taskBean.account,
	//'taskStatus':'${taskBean.taskStatus}',
	//'taskDetailNum':'${taskBean.taskDetailNum}',
	//'overallProgress':'${taskBean.overallProgress}',
	//'changeTimes':'${taskBean.changeTimes}',
	//'finishDate':taskBean.finishDate,
	//'startDate':taskBean.startDate,
	//'uniqueId':taskBean.uniqueId,
	//'deadLine':taskBean.deadLine,
	//'taskIconBean':"",
	//'detailList':"",
	//'token':token,
}
type NewUpdateTaskReq struct {
	Id string `json:"uniqueId,omitempty"`
}
