package entity

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey"`
	FirstName string     `gorm:"column:first_name;type:varchar(50)"`
	LastName  string     `gorm:"column:last_name;type:varchar(50)"`
	UserName  string     `gorm:"column:user_name;type:varchar(50)"`
	Email     string     `gorm:"column:email;type:varchar(50)"`
	Password  string     `gorm:"column:password;type:varchar(150)"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	CreatedBy string     `gorm:"column:created_by;type:varchar(50)"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedBy string     `gorm:"column:updated_by;type:varchar(50)"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy string     `gorm:"column:deleted_by;type:varchar(50)"`
}
