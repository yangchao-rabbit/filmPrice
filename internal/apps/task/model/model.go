package model

import (
	"filmPrice/internal/apps/task/dao"
	"filmPrice/internal/models"
)

type TaskListReq struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Filter   string `json:"filter" form:"filter"`
	Type     string `json:"type" form:"type"`
}

type TaskListResp struct {
	Rows  []*dao.TaskModel `json:"rows"`
	Total int64            `json:"total"`
}

type TaskCreateReq struct {
	Type     string           `json:"type" form:"type" binding:"required"`
	Name     string           `json:"name" form:"name" binding:"required"`
	Cron     string           `json:"cron" form:"cron"`
	FuncName string           `json:"func_name" form:"func_name" binding:"required"`
	Params   models.CustomMap `json:"params" form:"params" binding:"required"`
	IsActive bool             `json:"is_active" form:"is_active"`
	Desc     string           `json:"desc" form:"desc"`
}

type TaskUpdateReq struct {
	Type     string           `json:"type" form:"type"`
	Name     string           `json:"name" form:"name""`
	Cron     string           `json:"cron" form:"cron"`
	FuncName string           `json:"func_name" form:"func_name"`
	Params   models.CustomMap `json:"params" form:"params"`
	IsActive bool             `json:"is_active" form:"is_active"`
	Desc     string           `json:"desc" form:"desc"`
}

type IDReq struct {
	ID string `json:"id" form:"id" uri:"id" binding:"required"`
}

type TaskTestCronReq struct {
	Spec string `json:"spec" form:"spec" binding:"required"`
}
