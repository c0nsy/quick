// websocket/connection.go
package websocket

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	// The WebSocket connection
	WS *websocket.Conn

	// Buffered channel of outbound messages
	Send chan []byte
}

// WritePump writes messages from the hub to the WebSocket connection
func (c *Connection) WritePump() {
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.WS.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.WS.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// ReadPump reads messages from the WebSocket connection and sends them to the hub
func (c *Connection) ReadPump(hub *Hub, sessionID string) {
	defer func() {
		hub.Unregister <- &Subscription{Conn: c, SessionID: sessionID}
		c.WS.Close()
	}()
	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			break
		}
		hub.Broadcast <- Message{SessionID: sessionID, Data: message}
	}
}
