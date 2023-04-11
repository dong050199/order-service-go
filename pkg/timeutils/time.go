package timeutils

import "time"

func BeginningOfDay(datetime time.Time) time.Time {
	locationGMT07, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		panic(err)
	}
	y, m, d := datetime.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, locationGMT07)
}
