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
	userID := "d2c7ad7d-f29d-4f24-adb8-4e4650641588"
	t.Log(url)
	st, err := json.Marshal(struct {
		Method         string `json:"method"`
		IsAttended     int    `json:"is_attended"`
		InvitationType string `json:"invitation_type"`
	}{"POST", 1, "Wedding"})
	if err != nil {
		t.Fatal(err)
	}
	statusCode, err := RestClient(http.MethodPost, url, userID, st)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(statusCode)
}
