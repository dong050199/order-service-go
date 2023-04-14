package router

import (
	"order-service/service/http/handler"

	"github.com/gin-gonic/gin"
)

type ProductRouter interface {
	Register(rg *gin.RouterGroup)
}

type productRouter struct {
	productHandler handler.ProductHandler
}

func NewProductRouter(
	productHandler handler.ProductHandler,
) ProductRouter {
	return &productRouter{
		productHandler: productHandler,
	}
}

func (sr *productRouter) Register(r *gin.RouterGroup) {
	orderGroup := r.Group("/product")
	// orderGroup.Use(middleware.JwtAuthMiddleware())
	{
		orderGroup.GET("/list", sr.productHandler.GetList())
		orderGroup.GET("/:id", sr.productHandler.GetDetails())
	}
}
