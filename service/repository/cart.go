package repository

import (
	"order-service/pkg/infra"
	"order-service/service/model/entity"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type IcartRepo interface {
	GetUserCart(ctx context.Context, cartID uint) (cart entity.Cart, err error)
	UpdateCartUser(ctx context.Context, req []entity.ProductCart) error
	DeleteCartUser(ctx context.Context, req []entity.ProductCart) error
	CreateSaleOrder(ctx context.Context, req entity.Order) error
}

type cartRepo struct {
	ormDB *gorm.DB
}

func NewCartRepo() IcartRepo {
	return &cartRepo{
		ormDB: infra.GetDB(),
	}
}

func (c *cartRepo) GetUserCart(
	ctx context.Context,
	cartID uint,
) (cart entity.Cart, err error) {
	err = c.ormDB.Preload("ProductCart").Where("id = ?", cartID).First(&cart).Error
	return
}

func (r *cartRepo) UpdateCartUser(ctx context.Context, req []entity.ProductCart) error {
	return r.ormDB.Save(&req).Error
}

func (r *cartRepo) DeleteCartUser(ctx context.Context, req []entity.ProductCart) error {
	return r.ormDB.Delete(&req).Error
}

func (r *cartRepo) CreateSaleOrder(ctx context.Context, req entity.Order) error {
	return r.ormDB.Create(&req).Error
}
