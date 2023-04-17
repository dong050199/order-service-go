package repository

import (
	"order-service/pkg/infra"
	"order-service/service/model/entity"
	"order-service/service/model/request"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type IproductRepo interface {
	GetList(req request.PagingRequest) (meals []entity.Product, err error)
	GetByID(id uint) (meals entity.Product, err error)
	GetPaging(req request.PagingRequest) (totalPages int, err error)
	GetByIDs(req []uint) (resp []entity.Product, err error)
}

type productRepo struct {
	ormDB *gorm.DB
}

func NewProductRepo() IproductRepo {
	return &productRepo{
		ormDB: infra.GetDB(),
	}
}

func (r *productRepo) Insert(body entity.Product, tx *gorm.DB) error {
	return tx.Create(&body).Error
}

func (r *productRepo) Update(body entity.Product, tx *gorm.DB) error {
	return tx.Save(&body).Error
}

func (r *productRepo) Delete(body entity.Product, tx *gorm.DB) error {
	return tx.Save(&entity.Product{
		ID:        body.ID,
		DeletedAt: &timeNow,
		DeletedBy: body.DeletedBy,
	}).Error
}

func (m *productRepo) GetByID(
	id uint,
) (resp entity.Product, err error) {
	dbQuery := m.ormDB.Where("id = ?", id).Find(&resp)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}
	return
}

func (m *productRepo) GetList(
	req request.PagingRequest,
) (resp []entity.Product, err error) {
	dbQuery := m.ormDB.Offset(req.GetOffsetFromRequest()).Limit(req.Size).Order("id ASC").Find(&resp)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}
	return
}

func (m *productRepo) GetPaging(
	req request.PagingRequest,
) (totalPages int, err error) {
	var NumberRecode int64
	dbQuery := m.ormDB.Model(&entity.Product{}).Count(&NumberRecode)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}

	totalPages = int(NumberRecode) / int(req.Size)
	if totalPages == 0 {
		return
	}

	return
}

func (m *productRepo) GetByIDs(
	req []uint,
) (resp []entity.Product, err error) {
	if len(req) == 0 {
		return
	}
	dbQuery := m.ormDB.Find(&resp, req)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		if dbQuery.Error == gorm.ErrEmptySlice {
			err = nil
			return
		}
		return
	}
	return
}

func (p *productRepo) UpdateProductTx(
	ctx context.Context,
	tx *gorm.DB,
	req []entity.Product,
) (err error) {
	return tx.Save(&req).Error
}
