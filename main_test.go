package main

import (
	"HPE-golang-test/routes"
	"HPE-golang-test/services/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_serverSetup(t *testing.T) {
	server := routes.RouteSettings()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_reserveOK(t *testing.T) {
	// arrange
	server := routes.RouteSettings()
	w := httptest.NewRecorder()

	data := url.Values{
		"user_id":          {"U677970b5c9c674e88142e29027ac4ef2"},
		"reply_token":      {"861f9264aa6b4babb16041e8b9d0cec8"},
		"reserve_title":    {"test"},
		"reserve_datetime": {"2023-10-20T12:14"},
		"reserve_content":  {"testtestsssssssssssssssssssssssssssssststetsetsdvsdg"},
	}
	body := strings.NewReader(data.Encode())

	// act
	req, _ := http.NewRequest("POST", "/reserve", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	server.ServeHTTP(w, req)
	reserve, _ := reserveVerifyDBAndCleanUp(t, w.Body.String())

	// assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "New reserve done")
	assert.Contains(t, w.Body.String(), "reserve_id")
	assert.Equal(t, reserve.UserID, data.Get("user_id"))
	assert.Equal(t, reserve.ReplyToken, data.Get("reply_token"))
	assert.Equal(t, reserve.Content, data.Get("reserve_content"))
}

func Test_reserveContentTooLong(t *testing.T) {
	// arrange
	server := routes.RouteSettings()
	w := httptest.NewRecorder()

	data := url.Values{
		"user_id":          {"U677970b5c9c674e88142e29027ac4ef2"},
		"reply_token":      {"861f9264aa6b4babb16041e8b9d0cec8"},
		"reserve_title":    {"test"},
		"reserve_datetime": {"2023-10-20T12:14"},
		"reserve_content":  {"testtestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdgtesttestsssssssssssssssssssssssssssssststetsetsdvsdg"},
	}
	body := strings.NewReader(data.Encode())

	// act
	req, _ := http.NewRequest("POST", "/reserve", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	server.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "Error")
}

func reserveVerifyDBAndCleanUp(t *testing.T, jsonData string) (models.Reserve, error) {
	v := make(map[string]interface{})
	json.Unmarshal([]byte(jsonData), &v)
	// data := v.(map[string]string)

	reserveId := fmt.Sprintf("%v", v["reserve_id"])

	reserve, err := models.ReservationModel.QueryById(reserveId)

	//TODO: Clean up
	t.Cleanup(func() {
		models.ReservationModel.DeleteById(reserve.ID)
	})

	return reserve, err
}
