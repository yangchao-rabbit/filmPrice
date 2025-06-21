package model

type LocalLoginReq struct {
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

type LarkCallbackReq struct {
	Code  string `json:"code" form:"code"`
	State string `json:"state" form:"state"`
}
