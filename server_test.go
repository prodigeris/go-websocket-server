package main

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("Should return homepage", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		NewServer().ServeHTTP(response, request)

		assert.Equal(t, "Hi!", response.Body.String())
	})

	t.Run("Should open websockets on ws endpoint", func(t *testing.T) {

		server := httptest.NewServer(NewServer())
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err == nil {
			defer ws.Close()
		}

		assert.NoError(t, err)
	})
}
