package model

import "filmPrice/internal/apps/system/dao"

type PermListReq struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Filter   string `json:"filter" form:"filter"`
	Method   string `json:"method" form:"method"`
}

type PermListResp struct {
	Rows  []*dao.SystemPermModel `json:"rows"`
	Total int64                  `json:"total"`
}
