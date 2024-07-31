package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/min-tomato/online-shop/backend/internal/service"
	"github.com/min-tomato/online-shop/backend/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 2001, []string{"Bui", "Tuan", "Minh"})
}
