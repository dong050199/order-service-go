package entity

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	Proucts    []ProductOrder
	TotalPrice int        `gorm:"total_price"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	CreatedBy  string     `gorm:"column:created_by;type:varchar(50)"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	UpdatedBy  string     `gorm:"column:updated_by;type:varchar(50)"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
	DeletedBy  string     `gorm:"column:deleted_by;type:varchar(50)"`
}

type ProductOrder struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	OrderID   uint
	Quantity  int        `gorm:"column:quantity"`
	Price     float64    `gorm:"column:price"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	CreatedBy string     `gorm:"column:created_by;type:varchar(50)"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedBy string     `gorm:"column:updated_by;type:varchar(50)"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy string     `gorm:"column:deleted_by;type:varchar(50)"`
}
