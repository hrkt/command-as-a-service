package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/exec", strings.NewReader("some input"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	json := `{"result":"SOME INPUT"}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
