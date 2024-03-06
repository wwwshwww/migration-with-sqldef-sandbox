package mycontext

import "context"

type Context interface {
	Raw() context.Context
	SessionID() (string, bool)
	SessionToken() (string, bool)
}

type myContext struct {
	ctx context.Context
}

func From(ctx context.Context) Context {
	return &myContext{ctx: ctx}
}

func (c *myContext) Raw() context.Context {
	return c.ctx
}
func (c *myContext) SessionID() (string, bool) {
	si, ok := c.ctx.Value("sid").(string)
	return si, ok
}
func (c *myContext) SessionToken() (string, bool) {
	st, ok := c.ctx.Value("stoken").(string)
	return st, ok
}
