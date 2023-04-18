package response

import "order-service/service/model/entity"

type ListProductResponse struct {
	Products   []ProductResponse `json:"products"`
	TotalPage  int               `json:"total_pages"`
	Page       int               `json:"page"`
	TotalPrice int               `json:"total_price"`
}

type ProductResponse struct {
	entity.Product
	TotalPrice int `json:"total_price"`
}
