package request

type PagingRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (l *PagingRequest) GetOffsetFromRequest() int {
	offset := (l.Page - 1) * l.Size
	return offset
}
