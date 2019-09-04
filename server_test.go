package main

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var tenMiliseconds = 10 * time.Millisecond

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

	t.Run("Should respond to messages", func(t *testing.T) {

		server := httptest.NewServer(NewServer())
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err == nil {
			defer ws.Close()
		}

		sentData := []byte("How are you?")

		within(t, tenMiliseconds, func() {
			ws.WriteMessage(websocket.TextMessage, sentData)

			_, data, _ := ws.ReadMessage()
			assert.Equal(t, sentData, data)
		})
	})
}

func within(t *testing.T, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}
