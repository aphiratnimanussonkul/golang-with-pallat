package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

type fakeStore struct{}

func (fakeStore) NewTask(ttile string) error { return nil }

func TestTodoNewTaskHandler(t *testing.T) {
	payload := bytes.NewBuffer([]byte(`{"title":"task"}`))
	req := httptest.NewRequest("POST", "http://8080/todos", payload)
	w := httptest.NewRecorder()
	h := todoHandler{store: fakeStore{}}

	h.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Error("Expect 200 but got", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Error("Expect get content-type applicationjson but got", resp.Header.Get("Content-Type"))
	}
	fmt.Println(string(body))
}
