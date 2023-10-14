package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context

	spanctx context.Context // datadog span context.
}

func (c *Context) SpanCtx() context.Context {
	return c.spanctx
}

func New(c *gin.Context) *Context {
	if c == nil {
		return NewDefault()
	}
	return &Context{
		Context: c,
		spanctx: c.Request.Context(),
	}
}

func NewDefault() *Context {
	return &Context{
		Context: &gin.Context{},
		spanctx: context.Background(),
	}
}

func (c *Context) New(spanctx context.Context) *Context {
	return &Context{
		Context: c.Context,
		spanctx: spanctx,
	}
}

func (c *Context) Abstract() context.Context {
	return c
}

func Parse(ctx context.Context) (*Context, bool) {
	if c, ok := ctx.(*Context); ok {
		return c, true
	}
	return nil, false
}

func ChildCtx(ctx context.Context, spanctx context.Context) (context.Context, bool) {
	c, ok := Parse(ctx)
	if !ok {
		return spanctx, false
	}
	return c.New(spanctx), true
}
