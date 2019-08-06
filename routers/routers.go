package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"

	"test/controllers"
	"test/core"
	"test/middlewares"
	"test/pkg/validator"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	binding.Validator = new(validator.DefaultValidator)
	gin.SetMode(viper.GetString("mode"))

	router := gin.Default()
	router.Use(middlewares.RequestID())

	group := router.Group("/api")
	{
		userController := controllers.NewUserController()
		// 获取用户
		group.GET("/users/:id", core.Handle(userController.Get))
	}

	return router
}
