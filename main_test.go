package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_StatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	StatusHandler(res, req)

	exp := "OK\n"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
