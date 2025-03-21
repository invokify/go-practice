package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUseContext(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	r.Header.Set("User-ID", "123")

	handleRequest(w, r)

	response := w.Result()
	if response.StatusCode != http.StatusRequestTimeout {
		t.Errorf("Expected status code %d, got %d", http.StatusRequestTimeout, response.StatusCode)
	}
}
