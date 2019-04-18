package cmd

import (
	"context"
	"fmt"
)

type Context interface {
	context.Context
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	IsDebug() bool
}

type ctx struct {
	context.Context
}

func (c *ctx) IsDebug() bool {
	val := c.Value("debug")
	b, ok := val.(bool)
	if ok {
		return b
	}
	return false
}
func (c *ctx) Debugf(fmtStr string, vals ...interface{}) {
	if c.IsDebug() {
		fmt.Printf("[DEBUG] "+fmtStr+"\n", vals...)
	}
}

func (c *ctx) Infof(fmtStr string, vals ...interface{}) {
	if c.IsDebug() {
		fmt.Printf(fmtStr+"\n", vals...)
	}
}

// NewContext creates a new context and turns on debugging if debug is true
func NewContext(c context.Context, debug bool) Context {
	return &ctx{
		Context: context.WithValue(c, "debug", true),
	}
}
