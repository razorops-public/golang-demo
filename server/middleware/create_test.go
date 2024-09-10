package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"todo/middleware"
)

func TestGetTasks(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/task", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.GetAllTask)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected string = "null"
	if strings.Compare(rr.Body.String(), expected) == 0 {
		t.Errorf("handler returned unexpected body: got %T want %T",
			rr.Body.String(), "null")
	}
}

func TestCreateTask(t *testing.T) {
	body := []byte(`{"task":"hellow", "status":"pending"}`)
	req := httptest.NewRequest("POST", "/api/task", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.CreateTask)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
