package wrapper

import (
	"encoding/json"
	"net/http"
	"order-service/pkg/errormap"

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
	if message == nil {
		newmessage := errormap.ErrorMapMsg[statusCode]
		message = &newmessage
	}
	return &Response{
		Data:       data,
		Message:    message,
		Code:       errormap.ErrorMapCode[statusCode],
		StatusCode: statusCode,
	}
}

func NewSuccessResponse(c *gin.Context, data interface{}) *Response {
	return NewResponse(c, http.StatusOK, data, nil)
}

func JSONOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewSuccessResponse(c, data))
}
