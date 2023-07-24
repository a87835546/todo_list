package controllers

import (
	"github.com/gin-gonic/gin"
	"todo_list/internal/logic"
	"todo_list/internal/parameters"
)

type TaskCtl struct {
}

var tl *logic.TaskLogic

func (tc *TaskCtl) Create(ctx *gin.Context) {
	req := parameters.CreateReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespError(ctx, ParameterErrorCode, "解析参数异常")
	}
	task, err := tl.Create(&req)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, "db error")
	}
	RespOk(ctx, task)
}

func (tc *TaskCtl) Update(ctx *gin.Context) {
	RespOk(ctx, nil)
}

func (tc *TaskCtl) Delete(ctx *gin.Context) {
	RespOk(ctx, nil)
}
func (tc *TaskCtl) GetAllByUserId(ctx *gin.Context) {
	RespOk(ctx, nil)
}
