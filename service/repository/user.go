package repository

import (
	"order-service/pkg/infra"
	"order-service/pkg/timeutils"
	"order-service/service/model/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var timeNow = time.Now().UTC()

type IuserRepo interface {
	CreateUser(tx *gorm.DB, userInfo entity.User) (clientID uint, cartID uint, err error)
	ValidateUser(userInfo entity.User) (valid bool, userID uint, cartID uint, err error)
	UpdateUser(tx *gorm.DB, userInfo entity.User) error
	DeleteUser(tx *gorm.DB, userInfo entity.User) error
	GetUserInfoTopPage(tx *gorm.DB, userID uint) (userInfo entity.User, err error)
	GetMyPage(tx *gorm.DB, userID uint) (userInfo entity.User, err error)
}

type userRepo struct {
	ormDB *gorm.DB
}

func NewUserRepo() IuserRepo {
	return &userRepo{
		ormDB: infra.GetDB(),
	}
}

func (u *userRepo) GetUserInfoTopPage(tx *gorm.DB, userID uint) (userInfo entity.User, err error) {
	err = u.ormDB.
		Model(&entity.User{}).
		Preload("BodyRecords", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("created_at >= ?", timeNow.Add(-12*time.Hour))
		}).
		Preload("MealHistories", func(tx *gorm.DB) *gorm.DB {
			return tx.Limit(8).Order("created_at DESC")
		}).
		Where(&entity.User{ID: userID}).
		Find(&userInfo).Error

	return
}

func (u *userRepo) GetMyPage(tx *gorm.DB, userID uint) (userInfo entity.User, err error) {
	err = u.ormDB.
		Preload("BodyRecords", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("created_at >= ?", timeNow.Add(-12*time.Hour))
		}).
		Preload("ExerciseRecords", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("created_at >= ?", timeutils.BeginningOfDay(timeNow))
		}).
		Preload("DiaryRecords", func(tx *gorm.DB) *gorm.DB {
			return tx.Limit(10) // TODO: add filter param
		}).
		Where(&entity.User{ID: userID}).
		Find(&userInfo).Error

	return
}

func (u *userRepo) CreateUser(tx *gorm.DB, userInfo entity.User) (clientID uint, cartID uint, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	userInfo.Password = string(hashedPassword)
	userInfo.CreatedAt = &timeNow

	userInfo.Cart.CreatedAt = &timeNow
	dbExecute := tx.Create(&userInfo)
	if dbExecute.Error != nil {
		return 0, 0, dbExecute.Error
	}

	return userInfo.ID, userInfo.Cart.ID, nil
}

func (u *userRepo) ValidateUser(userInfo entity.User) (valid bool, userID uint, cartID uint, err error) {
	var userInfoDB entity.User
	dbQuery := u.ormDB.
		Preload("Cart").
		Where("deleted_at IS NULL").
		Where(&entity.User{UserName: userInfo.UserName, DeletedAt: nil}).
		Or(&entity.User{Email: userInfo.Email, DeletedAt: nil}).
		Find(&userInfoDB)
	if dbQuery.Error != nil {
		err = dbQuery.Error
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userInfoDB.Password), []byte(userInfo.Password)) == nil {
		valid = true
		userID = userInfo.ID
		cartID = userInfo.Cart.ID
		return
	}

	return
}

func (u *userRepo) UpdateUser(tx *gorm.DB, userInfo entity.User) error {
	dbQuery := tx.Save(&userInfo)
	if dbQuery.Error != nil {
		return dbQuery.Error
	}
	return nil
}

func (u *userRepo) DeleteUser(tx *gorm.DB, userInfo entity.User) error {
	dbQuery := tx.Save(&entity.User{
		ID:        userInfo.ID,
		DeletedBy: userInfo.DeletedBy,
		DeletedAt: &timeNow,
	})
	if dbQuery.Error != nil {
		return dbQuery.Error
	}
	return nil
}
