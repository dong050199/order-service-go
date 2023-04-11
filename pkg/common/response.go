package common

import (
	"encoding/json"
	"net/http"
	"order-service/pkg/ginutils/errors"
	"order-service/pkg/ginutils/tracking"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	Message    *string     `json:"message,omitempty"`
	TrackID    string      `json:"trace_id"`
}

func (r *Response) String() string {
	data, _ := json.Marshal(r) // nolint:errchkjson
	return string(data)
}

func NewResponse(c *gin.Context, statusCode int, data interface{}, message *string) *Response {
	return &Response{
		Data:       data,
		Message:    message,
		Code:       errors.ErrorMap[statusCode],
		StatusCode: statusCode,
		TrackID:    tracking.GetTrackIDFromContext(c),
	}
}

func NewSuccessResponse(c *gin.Context, data interface{}) *Response {
	return NewResponse(c, http.StatusOK, data, nil)
}

func NewErrorResponse(c *gin.Context, err AppError) *Response {
	return NewResponse(c, err.StatusCode, nil, &err.Message)
}

func JSONOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewSuccessResponse(c, data))
}
