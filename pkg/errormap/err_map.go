package errormap

import "net/http"

var (
	ErrorMapCode map[int]int
	ErrorMapMsg  map[int]string
)

func Initialize() error {
	return loadData()
}

// loadData loads data from database and save memcache
func loadData() error {
	ErrorMapCode = make(map[int]int)
	ErrorMapCode = map[int]int{
		http.StatusOK:                  Success,
		http.StatusBadRequest:          BadRequestErr,
		http.StatusInternalServerError: InternalServerError,
		http.StatusNotFound:            NotFound,
		http.StatusConflict:            ConflictError,
	}

	ErrorMapMsg = make(map[int]string)
	ErrorMapMsg = map[int]string{
		http.StatusOK:                  SuccessMess,
		http.StatusBadRequest:          BadRequestErrMess,
		http.StatusInternalServerError: InternalServerErrMess,
		http.StatusNotFound:            NotFoundErrMess,
	}
	return nil
}

const (
	SuccessMess           = "Thành công.!"
	ConflictErrMess       = "Thông tin đã tồn tại.!"
	BadRequestErrMess     = "Thông tin không hợp lệ.!"
	UnAuthorizedErrMess   = "Bạn có quyền truy xuất tài nguyên này.!"
	NotFoundErrMess       = "Thông tin không tồn tại.!"
	NotFetchDataErrMess   = "Không thể lấy dữ liệu.!"
	ForbiddenErrMess      = "Bạn không được phép truy xuất tài nguyên này.!"
	InternalServerErrMess = "Lỗi hệ thống.!"
)

const (
	// General Errors: 0 -> -49
	// Processing indicate success but the object is being processed
	Processing = 2
	// Success indicates no error
	Success = 1
	// Unknown error indicates unknown state or step
	Unknown = 0
	// BadRequest error
	BadRequestErr = -1
	// NotFound error
	NotFound = -2
	// AuthenFailed error
	AuthenticationFailed = -3
	// Internal server error
	InternalServerError = -4
	// IllegalStateError
	IllegalStateError = -5
	// SendMessageError
	SendMessageError = -6
	// Call Internal API Error
	CallInternalAPIError = -7
	// Invalid Data
	InvalidData = -8
	// SerializeError
	SerializingError = -9
	// DeserializeError
	DeserializingError = -10
	// CastingError
	CastingError = -11
	// ParsingError
	ParsingError = -12
	// ConflictError
	ConflictError = -13
	// Call GRPC Internal API Error
	CallGRPCAPIError = -14
	// Call Senpay Service Error
	CallSenPayServiceErr = 5
	// Get Senpay Token Error
	GetSenPayTokenError = 6
	// Get Senpay User Info Error
	GetSenPayUserInfoError = 7
	// Call Order Service Error
	CallOrderServiceError = 8
	// Order Status Invalid
	OrderStatusInvalid = 9
	// Payment Status Invalid
	PaymentStatusInvalid = 10
)
