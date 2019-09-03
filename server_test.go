package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("Should return homepage", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		WebSocketServer{}.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Hi!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
