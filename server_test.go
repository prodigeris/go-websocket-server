package main

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var timeout = 10 * time.Millisecond

func TestHttpServer(t *testing.T) {
	t.Run("Should serve HTML client", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		newServer().ServeHTTP(response, request)

		expected, _ := ioutil.ReadFile("client.html")

		assert.Equal(t, string(expected), response.Body.String())
	})
}

func TestWebSocketServer(t *testing.T) {

	type test struct {
		description     string
		sentMessage     []byte
		receivedMessage []byte
	}

	tests := []test{
		{description: "Should open WebSockets on /ws endpoint"},
		{description: "Should be able to accept message", sentMessage: []byte("How are you")},
		{
			description:     "Should be able to send message back",
			sentMessage:     []byte("How are you"),
			receivedMessage: []byte("How are you"),
		},
		{
			description:     "Should send transformed message back",
			sentMessage:     []byte("How are you?"),
			receivedMessage: []byte("How are you!"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			ws := startWebsocketServer(t)
			defer ws.Close()

			if tc.sentMessage != nil {
				assertWebSocketMessageSent(ws, tc.sentMessage, t)
			}

			if tc.receivedMessage != nil {
				within(t, timeout, func() {
					assertWebSocketReadMessageIsCorrect(ws, tc.receivedMessage, t)
				})
			}
		})
	}

	t.Run("Should close connection when non-text message received", func(t *testing.T) {
		ws := startWebsocketServer(t)

		err := ws.WriteMessage(websocket.BinaryMessage, []byte("Binary Message"))
		if err != nil {
			t.Fatal("Failed to send message to websocket")
		}

		within(t, timeout, func() {
			messageType, _, _ := ws.ReadMessage()

			assert.Equal(t, -1, messageType)
		})
	})
}

func startWebsocketServer(t *testing.T) *websocket.Conn {
	server := httptest.NewServer(newServer())
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
