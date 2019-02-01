package app

import "github.com/EDDYCJY/edge-pprof/pkg/e"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{
		Code: e.SUCCESS,
		Msg:  e.StatusText(e.SUCCESS),
		Data: nil,
	}
}

func (r *Response) Set(code int) {
	r.Code = code
	r.Msg = e.StatusText(code)
}
