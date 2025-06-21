package model

import "filmPrice/internal/apps/system/dao"

type UserListReq struct {
	Filter   string `json:"filter" form:"filter"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type UserListResp struct {
	Rows  []*dao.SystemUserModel `json:"rows"`
	Total int64                  `json:"total"`
}

type UserCreateReq struct {
	Type     string `json:"type" form:"type" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password"`
	Desc     string `json:"desc" form:"desc"`
}

type UserUpdateReq struct {
	Type     string `json:"type" form:"type"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Desc     string `json:"desc" form:"desc"`
}
