package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"test/core"
)

const (
	name = "X-Request-Id"
)

func RequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := &core.Context{Context: context}
		requestId := ctx.Request.Header.Get(name)
		if requestId == "" {
			requestId = uuid.NewV4().String()
			ctx.Writer.Header().Set(name, requestId)
		}
		ctx.Next()
	}
}
