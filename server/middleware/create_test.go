package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestGetTasks(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/task", nil)
	// if err != nil {
	// 	t.Fatal()
	// }

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllTask)
	handler.ServeHTTP(rr,req)
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
