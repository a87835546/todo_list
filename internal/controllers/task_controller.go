package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/parameters"
)

type TaskCtl struct {
}

var tl *logic.TaskLogic

func (tc *TaskCtl) Create(ctx *gin.Context) {
	req := parameters.CreateReq{}
	ParserReqParameters(&req, ctx)
	log.Printf("req---->>>> %v", req)
	task, err := tl.Create(&req)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, "db error")
	} else {
		RespOk(ctx, task)
	}
}

func (tc *TaskCtl) Update(ctx *gin.Context) {
	RespOk(ctx, nil)
}

func (tc *TaskCtl) Delete(ctx *gin.Context) {
	RespOk(ctx, nil)
}
func (tc *TaskCtl) GetAllByUserId(ctx *gin.Context) {
	mp := make(map[string]any, 0)
	err := ctx.BindJSON(&mp)
	id := mp["id"]

	tasks, err := tl.QueryByUserId(id)
	if err != nil {
		RespError(ctx, QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, tasks)
	}
}
