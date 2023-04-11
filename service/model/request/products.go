package request

type PagingRequest struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}
