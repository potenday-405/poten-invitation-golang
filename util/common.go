package util

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func StringToTime(t string) (*time.Time, error) {
	if len(t) == 0 {
		return nil, nil
	}
	if len(t) != 12 {
		return nil, errors.New("invalid time parameter")
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(t) {
		return nil, errors.New("invalid time parameter")
	}
	year, _ := strconv.Atoi(t[0:4])
	month, _ := strconv.Atoi(t[4:6])
	day, _ := strconv.Atoi(t[6:8])
	hour, _ := strconv.Atoi(t[8:10])
	minute, _ := strconv.Atoi(t[10:12])

	// time.Time 객체를 생성합니다.
	date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
	return &date, nil
}
