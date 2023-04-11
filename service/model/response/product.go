package response

import "order-service/service/model/entity"

type ListProductResponse struct {
	Products  []entity.Product `json:"products"`
	TotalPage int              `json:"total_pages"`
	Page      int              `json:"page"`
}
