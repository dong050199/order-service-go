package entity

import "time"

type Cart struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	ProductCart []ProductCart
	CreatedAt   *time.Time `gorm:"column:created_at"`
	CreatedBy   string     `gorm:"column:created_by;type:varchar(50)"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	UpdatedBy   string     `gorm:"column:updated_by;type:varchar(50)"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
	DeletedBy   string     `gorm:"column:deleted_by;type:varchar(50)"`
}
type ProductCart struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	Quantity  int `gorm:"column:quantity"`
	ProductID uint
}
