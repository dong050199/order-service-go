package handler

import (
	"net/http"
	"order-service/pkg/middleware"
	"order-service/pkg/tracing"
	"order-service/service/model/request"
	"order-service/service/usecase"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase usecase.IcartUsercase
}

func NewCartHandler(
	cartUsecase usecase.IcartUsercase,
) CartHandler {
	return CartHandler{
		cartUsecase: cartUsecase,
	}
}

// @Summary Get cart
// @Description Get cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} response.ListProductResponse
// @Failure 400 {object} response.ListProductResponse
// @Failure 500 {object} response.ListProductResponse
// @Router /cart [get]
func (h *CartHandler) GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracing.StartSpanFromCtx(c, "GetCart")
		defer span.Finish()
		data, err := h.cartUsecase.GetCart(ctx, uint(middleware.ExtractCartIDFromContext(c)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// @Summary Update cart
// @Description Update cart
// @Tags Cart
// @Accept json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param model body request.UpdateCartRequest true "model"
// @Success 200 {object} response.ListProductResponse
// @Failure 400 {object} response.ListProductResponse
// @Failure 500 {object} response.ListProductResponse
// @Router /cart [put]
func (h *CartHandler) UpdateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracing.StartSpanFromCtx(c, "GetCart")
		defer span.Finish()

		var req request.UpdateCartRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err := h.cartUsecase.UpdateCart(ctx, uint(middleware.ExtractCartIDFromContext(c)), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Create order
// @Description Create order
// @Tags Order
// @Accept json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} response.ListProductResponse
// @Failure 400 {object} response.ListProductResponse
// @Failure 500 {object} response.ListProductResponse
// @Router /cart/create-order [get]
func (h *CartHandler) CreateSaleOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracing.StartSpanFromCtx(c, "GetCart")
		defer span.Finish()

		err := h.cartUsecase.CreateSalesOrder(ctx,
			uint(middleware.ExtractCartIDFromContext(c)),
			uint(middleware.ExtractUserIDFromContext(c)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
