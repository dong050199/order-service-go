package repository

import (
	"order-service/pkg/infra"
	"order-service/service/model/entity"
	"order-service/service/model/request"

	"gorm.io/gorm"
)

type IproductRepo interface {
	GetList(req request.PagingRequest) (meals []entity.Product, err error)
	GetByID(id uint) (meals entity.Product, err error)
	GetPaging(req request.PagingRequest) (totalPages int, err error)
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
	dbQuery := m.ormDB.Offset(req.Offset).Limit(req.Limit).Order("id ASC").Find(&resp)
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

	totalPages = int(NumberRecode) / int(req.Limit)
	if totalPages == 0 {
		return
	}

	return
}
