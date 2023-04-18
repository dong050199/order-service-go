package usecase

import (
	"context"
	"order-service/pkg/tracing"
	"order-service/service/model/entity"
	"order-service/service/model/request"
	"order-service/service/model/response"
	"order-service/service/repository"

	"github.com/labstack/gommon/log"
)

type IproductUsecase interface {
	GetList(ctx context.Context, req request.PagingRequest) (resp response.ListProductResponse, err error)
	GetDetails(ctx context.Context, productID uint) (resp entity.Product, err error)
}

type productUsecase struct {
	productRepo repository.IproductRepo
}

func NewProductUsecase(
	productRepo repository.IproductRepo,
) IproductUsecase {
	return &productUsecase{productRepo}
}

func (p *productUsecase) GetList(
	ctx context.Context,
	req request.PagingRequest,
) (resp response.ListProductResponse, err error) {
	totalPage, err := p.productRepo.GetPaging(req)
	if err != nil {
		log.Errorf("Get List Paging error: %v", err)
		return
	}

	campaigns, err := p.productRepo.GetList(req)
	if err != nil {
		log.Errorf("Get List Paging error: %v", err)
		return
	}

	resp.Page = req.Page
	resp.TotalPage = totalPage

	for _, product := range campaigns {
		resp.Products = append(resp.Products, response.ProductResponse{
			Product:    product,
			TotalPrice: 0,
		})
	}

	return
}

func (p *productUsecase) GetDetails(
	ctx context.Context,
	productID uint,
) (resp entity.Product, err error) {
	span, _ := tracing.StartSpanFromCtx(ctx, "GetDetails")
	defer span.Finish()
	resp, err = p.productRepo.GetByID(productID)
	if err != nil {
		log.Errorf("Get List Paging error: %v", err)
		return
	}
	return
}
