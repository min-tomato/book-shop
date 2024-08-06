package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/min-tomato/online-shop/backend/internal/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}

	return r
}
