package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRequestInfo(t *testing.T) {
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Hello", "world")
	if err != nil {
		t.Fatal(err)
	}

	RequestInfo(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: go %v want %v", status, http.StatusOK)
	}

	resp := rec.Body.String()

	expectedTimeFormat := time.DateTime
	expectedParts := []string{
		fmt.Sprintf("%s\n", time.Now().Format(expectedTimeFormat)),
		"GET / HTTP/1.1\n",
		"Remote address: ",
		"Host: ",
		"Headers:",
		"Hello: world\n",
	}

	for _, part := range expectedParts {
		if !strings.Contains(resp, part) {
			t.Errorf("Response body missing expected part: %s", part)
		}
	}
}
