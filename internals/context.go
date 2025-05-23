package server

import (
	"http-server/internals/request"
	"http-server/internals/response"
)

type Context struct {
	Req *request.Request
	Res *response.Response
}

func NewContext(req *request.Request, res *response.Response) *Context {
	return &Context{Req: req, Res: res}
}
