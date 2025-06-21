package model

import "filmPrice/internal/apps/system/dao"

type GroupListReq struct {
	Filter   string `json:"filter" form:"filter"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type GroupListResp struct {
	Rows  []*dao.SystemGroupModel `json:"rows"`
	Total int64                   `json:"total"`
}

type GroupCreateReq struct {
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc"`
}

type GroupUpdateReq struct {
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc"`
}
