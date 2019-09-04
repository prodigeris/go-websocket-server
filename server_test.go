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

var timeout = 10 * time.Millisecond

func TestServer(t *testing.T) {
	t.Run("Should return homepage", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		NewServer().ServeHTTP(response, request)

		assert.Equal(t, "Hi!", response.Body.String())
	})

	t.Run("Should open websockets on ws endpoint", func(t *testing.T) {

		ws := startWebsocketServer(t)
		defer ws.Close()
	})

	t.Run("Should be able to accept message", func(t *testing.T) {

		ws := startWebsocketServer(t)
		defer ws.Close()

		sentData := []byte("How are you")
		assertWebSocketMessageSent(ws, sentData, t)
	})

	t.Run("Should be receive message", func(t *testing.T) {

		ws := startWebsocketServer(t)
		defer ws.Close()

		sentData := []byte("How are you")
		assertWebSocketMessageSent(ws, sentData, t)

		within(t, timeout, func() {
			assertWebSocketReadMessageIsCorrect(ws, sentData, t)
		})
	})
}

func startWebsocketServer(t *testing.T) *websocket.Conn {
	server := httptest.NewServer(NewServer())
	defer server.Close()
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatal("Cannot start a websocket server")
	}

	return ws
}

func assertWebSocketMessageSent(ws *websocket.Conn, sentData []byte, t *testing.T) {
	err := ws.WriteMessage(websocket.TextMessage, sentData)
	if err != nil {
		t.Fatal("Failed to send message to websocket")
	}
}

func assertWebSocketReadMessageIsCorrect(ws *websocket.Conn, expectedData []byte, t *testing.T) {
	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatal("Failed to send message to websocket")
	}

	assert.Equal(t, expectedData, message)
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
