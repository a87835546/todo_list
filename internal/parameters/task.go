package parameters

type CreateReq struct {
	TaskName        string `json:"taskName"`
	TaskType        string `json:"taskType"`
	Account         string `json:"account"`
	TaskStatus      int    `json:"taskStatus"`
	TaskDetailNum   int    `json:"taskDetailNum"`
	OverallProgress int    `json:"overallProgress"`
	ChangeTimes     string `json:"changeTimes"`
	FinishDate      string `json:"finishDate"`
	StartDate       string `json:"startDate"`
	DeadLine        string `json:"deadLine"`
	TaskIconBean    string `json:"taskIconBean"`
	DetailList      string `json:"detailList"`
	Token           string `json:"token"`
}
