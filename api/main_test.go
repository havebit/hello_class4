package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeStore struct{}

func (fakeStore) NewTask(title string) error { return nil }

func TestTodoNewTaskHandler(t *testing.T) {
	payload := bytes.NewBuffer([]byte(`{"title":"test"}`))

	req := httptest.NewRequest("POST", "http://example.com/foo", payload)
	w := httptest.NewRecorder()
	h := todoHandler{store: fakeStore{}}
	h.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Error("expected status 200 OK, got", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Error("expected Content-Type application/json, got", resp.Header.Get("Content-Type"))
	}
	fmt.Println(string(body))
}
