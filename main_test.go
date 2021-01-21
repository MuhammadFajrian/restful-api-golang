package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HOme(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	expectedResponse := `{"message":"Hello World", "status":"200"}`
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.Code, "Invalid Response Code")
	assert.Equal(t, expectedResponse, string(bytes.TrimSpace(responseBody)))
}
