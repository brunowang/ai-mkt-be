package service

import "context"

type RspContext struct {
	Ctx context.Context
	Rsp interface{}
}

func NewRspContext(ctx context.Context, rsp interface{}) *RspContext {
	return &RspContext{
		Ctx: ctx,
		Rsp: rsp,
	}
}

type ErrContext struct {
	Ctx context.Context
	Err error
}

func NewErrContext(ctx context.Context, err error) *ErrContext {
	return &ErrContext{
		Ctx: ctx,
		Err: err,
	}
}

func (c *ErrContext) Error() string {
	if c == nil {
		return ""
	}
	if c.Err == nil {
		return ""
	}
	return c.Err.Error()
}
