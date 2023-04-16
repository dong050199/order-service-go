package router

import (
	"order-service/pkg/middleware"
	"order-service/service/http/handler"

	"github.com/gin-gonic/gin"
)

type CartRouter interface {
	Register(rg *gin.RouterGroup)
}

type cartRouter struct {
	cartHandler handler.CartHandler
}

func NewCartRouter(
	cartHandler handler.CartHandler,
) CartRouter {
	return &cartRouter{
		cartHandler: cartHandler,
	}
}

func (sr *cartRouter) Register(r *gin.RouterGroup) {
	page := r.Group("/cart")
	page.Use(middleware.JwtAuthMiddleware())
	{
		page.GET("", sr.cartHandler.GetCart())
		page.POST("", sr.cartHandler.UpdateCart())
		page.GET("/create-order", sr.cartHandler.CreateSaleOrder())
	}
}
