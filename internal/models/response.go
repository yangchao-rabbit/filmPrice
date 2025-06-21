package models

import "encoding/json"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func FailResp(code int, msg string) *Response {
	return NewResponse(code, msg, nil)
}

func SuccessResp(data interface{}) *Response {
	return NewResponse(0, "success", data)
}

func (r *Response) String() string {
	m, _ := json.Marshal(r)
	return string(m)
}
