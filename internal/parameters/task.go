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
