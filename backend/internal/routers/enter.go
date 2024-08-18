package routers

import "github.com/min-tomato/online-shop/backend/internal/routers/user"

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
