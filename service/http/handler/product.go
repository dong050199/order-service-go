package handler

import (
	"fmt"
	"net/http"
	"order-service/pkg/tracing"
	"order-service/service/model/request"
	"order-service/service/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ProductHandler struct {
	productUsecase usecase.IproductUsecase
}

func NewProductHandler(
	productUsecase usecase.IproductUsecase,
) ProductHandler {
	return ProductHandler{
		productUsecase: productUsecase,
	}
}

func composeConfigHandlerName(name string) string {
	return fmt.Sprintf("config_handler_%s", name)
}

// Product godoc
// @Summary Get details product
// @Description Get details product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "campaign_id"
// @Success 200 {object} entity.Product
// @Failure 400 {object} entity.Product
// @Failure 500 {object} entity.Product
// @Router /product/{id} [get]
func (h *ProductHandler) GetDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracing.StartSpanFromCtx(c, "GetDetails")
		defer span.Finish()
		data, err := h.productUsecase.GetDetails(ctx, cast.ToUint(c.Param("id")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// @Summary Get details product
// @Description Get details product
// @Tags Product
// @Accept json
// @Produce json
// @Param offset query string true "offset"
// @Param limit query string true "limit"
// @Param total query string true "total"
// @Success 200 {object} response.ListProductResponse
// @Failure 400 {object} response.ListProductResponse
// @Failure 500 {object} response.ListProductResponse
// @Router /product/list [get]
func (h *ProductHandler) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := tracing.StartSpanFromCtx(c, "GetList")
		defer span.Finish()
		req := request.PagingRequest{
			Offset:     cast.ToInt(c.Query("offet")),
			Limit:      cast.ToInt(c.Query("limit")),
			TotalPages: cast.ToInt(c.Query("total")),
		}

		data, err := h.productUsecase.GetList(ctx, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}
