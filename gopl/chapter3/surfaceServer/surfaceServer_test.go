package main

import "testing"
import "net/http/httptest"

func TestHandlerContentType(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, r)

	contentType := w.Header().Get("Content-Type")
	if contentType != "image/svg+xml" {
		t.Error("Expected xml Content-Type, but got", contentType)
	}
}
