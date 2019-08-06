package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"test/core"
	"test/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: &services.UserService{},
	}
}

func (controller *UserController) Get(ctx *core.Context) {
	id := ctx.Param("id")

	user, err := controller.userService.GetById(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Fail(http.StatusNotFound, errors.New("用户不存在"))
			return
		}
		ctx.Fail(http.StatusInternalServerError, err)
		return
	}
	ctx.Success(http.StatusOK, user)
}
