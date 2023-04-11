package ginutils

import (
	"order-service/pkg/common"
	"order-service/pkg/ginutils/constants"
	"order-service/pkg/logger"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoverPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			dataError := map[string]interface{}{
				"method": c.Request.Method,
				"path":   c.Request.URL.EscapedPath(),
				"err":    err,
				"stack":  string(debug.Stack()),
			}
			logger.NewLogger().WithKeyword(c.Request.Context(), "gin-ErrRecover").WithOutput(dataError).Error()
			common.HandleError(c, common.NewInternalServerError(nil, constants.InternalServerErrMess))
		}
	}()

	c.Next()
}
