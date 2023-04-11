package logger

import (
	"context"
	"fmt"
	"order-service/pkg/ginutils/constants"
	"order-service/pkg/ginutils/timeutils"
	"time"

	"github.com/spf13/cast"
)

func WriteLogger(ctx context.Context, dataLog map[string]interface{}, name string, startTime time.Time) {
	allTime := cast.ToFloat64(timeutils.Since(startTime).Milliseconds())
	dataLog["allTime"] = allTime
	logger := NewLogger()
	if v, ok := dataLog[constants.ErrorKey]; ok {
		logger.WithKeyword(ctx, name+"_error").
			WithOutput(dataLog).
			WithResponseTime(allTime).
			WithErrorStr(fmt.Sprintf("%+v", v)).
			Error()
	} else if v, ok := dataLog[constants.WarnKey]; ok { // nolint: gocritic
		logger.WithKeyword(ctx, name+"_warn").
			WithOutput(dataLog).
			WithResponseTime(allTime).
			WithErrorStr(fmt.Sprintf("%+v", v)).
			Warn()
	} else if allTime > constants.WarnTime {
		logger.WithKeyword(ctx, name+"_warn").
			WithOutput(dataLog).
			WithResponseTime(allTime).
			Warn()
	} else {
		logger.WithKeyword(ctx, name+"_info").
			WithOutput(dataLog).
			WithResponseTime(allTime).
			Info()
	}
}
