// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/min-tomato/online-shop/backend/internal/controller"
	"github.com/min-tomato/online-shop/backend/internal/repo"
	"github.com/min-tomato/online-shop/backend/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
