package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/models"
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
	req := parameters.UpdateTaskReq{}
	ParserReqParameters(&req, ctx)
	err := tl.Update(&req)
	if err != nil {
		RespError(ctx, UpdateDBErrorCode, "db error")
	} else {
		RespOk(ctx, nil)
	}
}

func (tc *TaskCtl) Delete(ctx *gin.Context) {
	req := parameters.DeleteReq{}
	ParserReqParameters(&req, ctx)
	err := tl.Delete(&req)
	if err != nil {
		RespError(ctx, DeleteDBErrorCode, "db error")
	} else {
		RespOk(ctx, nil)
	}
}
func (tc *TaskCtl) GetAllByUserId(ctx *gin.Context) {
	mp := make(map[string]any)
	err := ctx.BindJSON(&mp)
	id := mp["id"]

	tasks, err := tl.QueryByUserId(id)
	if err != nil {
		RespError(ctx, QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, tasks)
	}
}
func (tc *TaskCtl) AddNewTask(ctx *gin.Context) {
	req := models.TaskMode{}
	ParserReqParameters(&req, ctx)
	log.Printf("req---->>>> %v", req)
	task, err := tl.CreateNew(&req)
	if err != nil {
		RespError(ctx, InsertDBErrorCode, "db error")
	} else {
		RespOk(ctx, task)
	}
}
func (tc *TaskCtl) UpdateNew(ctx *gin.Context) {
	req := parameters.NewUpdateTaskReq{}
	ParserReqParameters(&req, ctx)
	err := tl.NewUpdate(&req)
	if err != nil {
		RespError(ctx, UpdateDBErrorCode, "db error")
	} else {
		RespOk(ctx, nil)
	}
}
