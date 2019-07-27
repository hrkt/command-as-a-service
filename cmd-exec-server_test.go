package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSortWithOption(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/usr/bin/sort?-rn", strings.NewReader("a\nb"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	myServer := MyServer()
	myServer.ServeHTTP(w, req)

	result := "b\na\n"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, result, w.Body.String())
}

func TestExecuteSortWithoutOption(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/usr/bin/sort", strings.NewReader("a\nb"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	myServer := MyServer()
	myServer.ServeHTTP(w, req)

	result := "a\nb\n"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, result, w.Body.String())
}
