package router

import (
	"order-service/service/http/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Register(rg *gin.RouterGroup)
}

type userRouter struct {
	userHandlerr handler.UserHandler
}

func NewUserRouter(
	userHandlerr handler.UserHandler,
) UserRouter {
	return &userRouter{
		userHandlerr: userHandlerr,
	}
}

func (sr *userRouter) Register(r *gin.RouterGroup) {
	page := r.Group("/user")
	{
		page.POST("/register", sr.userHandlerr.Register())
		page.POST("/login", sr.userHandlerr.Login())
	}
}
