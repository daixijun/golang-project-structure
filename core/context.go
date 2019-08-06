package core

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	validator2 "test/pkg/validator"
)

type Context struct {
	*gin.Context
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Fail 失败返回
func (c *Context) Fail(statusCode int, data interface{}) {
	resp := &ErrorResponse{}

	if validationErrors, ok := data.(validator.ValidationErrors); ok {
		resp.Message = "参数错误"
		resp.Errors = validator2.GetErrorMessages(validationErrors)
	} else if err, ok := data.(error); ok {
		resp.Message = err.Error()
	} else if s, ok := data.(string); ok {
		resp.Message = s
	} else if err, ok := data.(*ErrorResponse); ok {
		resp = err
	}

	c.AbortWithStatusJSON(statusCode, resp)
}

// Success 成功返回
func (c *Context) Success(statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
