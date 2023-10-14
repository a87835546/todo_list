package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"todo_list/internal/logic"
	"todo_list/internal/models"
)

type TaskGroupController struct {
	ts *logic.TaskGroupService
}

func NewTaskGroupController() *TaskGroupController {
	return &TaskGroupController{
		ts: logic.NewTaskGroupService(),
	}
}

func (tc *TaskGroupController) QueryByUserId(ctx *gin.Context) {
	uid := ctx.Query("user_id")
	n, err := strconv.Atoi(uid)
	if err != nil {
		RespErrorWithMsg(ctx, ParameterErrorCode, "解析参数异常"+err.Error(), nil)
	} else {
		users, err := tc.ts.QueryByUserId(n)
		if err != nil {
			RespError(ctx, InsertDBErrorCode, err.Error())
		} else {
			RespOk(ctx, users)
		}
	}
}

func (tc *TaskGroupController) Add(ctx *gin.Context) {
	req := models.TaskGroupModel{}
	err := ParserReq(&req, ctx)
	if err != nil {
		RespErrorWithMsg(ctx, ParameterErrorCode, "解析参数异常"+err.Error(), nil)
	} else {
		log.Printf("req--->>> %#v", req)
		err = tc.ts.Add(&req)
		if err != nil {
			RespError(ctx, InsertDBErrorCode, err.Error())
		} else {
			RespOk(ctx, nil)
		}
	}
}

func (tc *TaskGroupController) Update(ctx *gin.Context) {
	req := models.TaskGroupModel{}
	err := ParserReq(&req, ctx)
	if err != nil {
		RespErrorWithMsg(ctx, ParameterErrorCode, "解析参数异常"+err.Error(), nil)
	} else {
		log.Printf("req--->>> %#v", req)
		err = tc.ts.Update(&req)
		if err != nil {
			RespError(ctx, InsertDBErrorCode, err.Error())
		} else {
			RespOk(ctx, nil)
		}
	}
}
