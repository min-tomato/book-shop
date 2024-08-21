//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/min-tomato/online-shop/backend/internal/controller"
	"github.com/min-tomato/online-shop/backend/internal/repo"
	"github.com/min-tomato/online-shop/backend/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
