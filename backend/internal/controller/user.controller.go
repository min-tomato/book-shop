package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/min-tomato/online-shop/backend/internal/services"
	"github.com/min-tomato/online-shop/backend/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
