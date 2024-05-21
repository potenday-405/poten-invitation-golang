package util

import (
	"strconv"
	"time"
)

func StringToTime(t string) time.Time {
	year, _ := strconv.Atoi(t[0:4])
	month, _ := strconv.Atoi(t[4:6])
	day, _ := strconv.Atoi(t[6:8])
	hour, _ := strconv.Atoi(t[8:10])
	minute, _ := strconv.Atoi(t[10:12])

	// time.Time 객체를 생성합니다.
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
}
