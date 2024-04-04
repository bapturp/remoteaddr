package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlerRequestInfo_DefaultResponse(t *testing.T) {
	log.SetOutput(bytes.NewBuffer(nil))

	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Hello", "world")
	if err != nil {
		t.Fatal(err)
	}

	HandlerRequestInfo(rec, req)
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

func TestHandlerRequestInfo_JSONResponse(t *testing.T) {
	log.SetOutput(bytes.NewBuffer(nil))

	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/?r=json", nil)
	req.Header.Set("Hello", "world")
	if err != nil {
		t.Fatal(err)
	}

	HandlerRequestInfo(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: go %v want %v", status, http.StatusOK)
	}

	resp := rec.Body.String()
	ri := &RequestInfo{}
	err = json.Unmarshal([]byte(resp), ri)
	if err != nil {
		t.Errorf("Response JSON malformatted: %s", err)
	}

	if ri.Time.IsZero() {
		t.Errorf("Expected Time field to be present in response")
	}
	if ri.Method != "GET" {
		t.Errorf("Expected Method field to be GET, got %s", ri.Method)
	}
	if !strings.Contains(ri.Proto, "HTTP") {
		t.Errorf("Expected Proto field to contain HTTP, got %s", ri.Proto)
	}
	if ri.Headers["Hello"] != "world" {
		t.Errorf("Expected Header Hello field should be world, got %s", ri.Headers["Hello"])
	}
}
