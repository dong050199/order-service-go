package errors

import "net/http"

var (
	ErrorMap map[int]int
)

func Initialize() error {
	return loadData()
}

// loadData loads data from database and save memcache
func loadData() error {
	ErrorMap = make(map[int]int)
	ErrorMap = map[int]int{
		http.StatusOK:                  Success,
		http.StatusBadRequest:          BadRequestErr,
		http.StatusInternalServerError: InternalServerError,
		http.StatusNotFound:            NotFound,
	}
	return nil
}
