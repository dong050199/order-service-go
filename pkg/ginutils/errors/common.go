package errors

import (
	"errors"
	"order-service/pkg/ginutils/constants"
)

var (
	ErrBadRequest   = errors.New(constants.BadRequestErrMess)
	ErrNotFound     = errors.New(constants.NotFoundErrMess)
	ErrInternal     = errors.New(constants.InternalServerErrMess)
	ErrRedis        = errors.New(constants.InternalServerErrMess)
	ErrNotExistPath = errors.New(constants.NotExistPathErrMess)
)
