package cmd

import (
	"context"
	"fmt"
)

type Context interface {
	context.Context
	Debugf(string, ...interface{})
}

type ctx struct {
	context.Context
}

func (c *ctx) Debugf(fmtStr string, vals ...interface{}) {
	val := c.Value("debug")
	b, ok := val.(bool)
	if ok && b {
		fmt.Printf(fmtStr+"\n", vals...)
	}
}

func NewContext(c context.Context, debug bool) Context {
	return &ctx{
		Context: context.WithValue(c, "debug", true),
	}
}
