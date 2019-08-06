package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context)

// Handle transform HandlerFuc to gin.HandlerFunc
func Handle(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}
		handler(ctx)
	}
}

func Ping(c *Context) {
	c.Success(http.StatusOK, "pong")
}

// NoRouterController 路径不存在
func NoRouterController(c *Context) {
	c.Fail(http.StatusNotFound, "您访问的资源去撸猫了...")
}

// NoMethodController method 不被允许
func NoMethodController(c *Context) {
	c.Fail(http.StatusMethodNotAllowed, "不允许的method")
}
