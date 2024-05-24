package util

import (
	"errors"
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
	"strconv"
	"time"
)

type UserScore struct {
	Method         string `json:"method"`
	IsAttended     int    `json:"is_attended"`
	InvitationType string `json:"invitation_type"`
}

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

func RestClient(method, url, accessToken string, body []byte) (int, error) {
	// fasthttp.Client 생성
	client := &fasthttp.Client{}

	// POST 요청 생성
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.Header.Set("access-token", accessToken)
	req.SetBody(body)

	// 요청 실행
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode(), nil
}
