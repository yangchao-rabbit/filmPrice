package model

type IDReq struct {
	ID string `json:"id" form:"id" uri:"id" binding:"required"`
}
