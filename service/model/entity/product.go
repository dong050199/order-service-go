package entity

import "time"

type Product struct {
	ID               uint       `json:"id"`
	CategoryID       int        `json:"category_id"`
	Price            float64    `json:"price"`
	Quantity         int        `json:"quantity"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	ShortDescription string     `json:"short_description"`
	ImageURL         string     `json:"image_url"`
	CreatedAt        *time.Time `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        string     `json:"updated_by"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DeletedBy        string     `json:"deleted_by"`
}

type Category struct {
	ID               uint       `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	ShortDescription string     `json:"short_description"`
	CreatedAt        *time.Time `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UpdatedBy        string     `json:"updated_by"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DeletedBy        string     `json:"deleted_by"`
	Products         []Product
}
