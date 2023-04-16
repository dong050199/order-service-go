package request

type UpdateCartRequest struct {
	Products []ProductCart `json:"produc_ids"`
}

type ProductCart struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
