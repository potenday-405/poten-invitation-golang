package util

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

func testInitializer() {
	if err := godotenv.Load("./../env/.env"); err != nil {
		log.Fatal(err)
	}
}

func TestStringToTime(t *testing.T) {
	_, err := StringToTime("20241200")
	assert.Error(t, err)
	_, err = StringToTime("202412001111")
	assert.NoError(t, err)
	_, err = StringToTime("20241200111a")
	assert.Error(t, err)
}

func TestRestClient(t *testing.T) {
	testInitializer()
	url := "http://" + os.Getenv("USER_SERVER") + "/user/score"
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkMmM3YWQ3ZC1mMjlkLTRmMjQtYWRiOC00ZTQ2NTA2NDE1ODgiLCJleHAiOjE3MTY1NDM0NzZ9.7mlugK4Wdo9J6GZdSB-rMzEAv4wmiQylAZdp7M0piEw"
	t.Log(url)
	st, err := json.Marshal(struct {
		Method     string `json:"method"`
		IsAttended int    `json:"is_attended"`
	}{"POST", 1})
	if err != nil {
		t.Fatal(err)
	}
	statusCode, err := RestClient(http.MethodPost, url, accessToken, st)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(statusCode)
}
