// websocket/handler.go
package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("sessionId") // Get session ID from query parameters
	if sessionID == "" {
		http.Error(w, "Missing session ID", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	conn := &Connection{WS: ws, Send: make(chan []byte, 256)}
	subscription := &Subscription{Conn: conn, SessionID: sessionID}

	hub.Register <- subscription

	go conn.WritePump()           // Handle outgoing messages
	conn.ReadPump(hub, sessionID) // Handle incoming messages
}
