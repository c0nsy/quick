// websocket/hub.go
package websocket

type Hub struct {
	// Registered connections by session ID
	Sessions map[string]map[*Connection]bool

	// Inbound messages from the connections
	Broadcast chan Message

	// Register requests from the connections
	Register chan *Subscription

	// Unregister requests from connections
	Unregister chan *Subscription
}

// Message represents a message broadcast to a room
type Message struct {
	SessionID string
	Data      []byte
}

// Subscription is a connection to a specific session
type Subscription struct {
	Conn      *Connection
	SessionID string
}

// NewHub initializes a new WebSocket hub
func NewHub() *Hub {
	return &Hub{
		Sessions:   make(map[string]map[*Connection]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *Subscription),
		Unregister: make(chan *Subscription),
	}
}

// Run listens for events like registering/unregistering connections and broadcasting messages
func (h *Hub) Run() {
	for {
		select {
		case subscription := <-h.Register:
			connections := h.Sessions[subscription.SessionID]
			if connections == nil {
				connections = make(map[*Connection]bool)
				h.Sessions[subscription.SessionID] = connections
			}
			h.Sessions[subscription.SessionID][subscription.Conn] = true

		case subscription := <-h.Unregister:
			connections := h.Sessions[subscription.SessionID]
			if connections != nil {
				if _, ok := connections[subscription.Conn]; ok {
					delete(connections, subscription.Conn)
					close(subscription.Conn.Send)
					if len(connections) == 0 {
						delete(h.Sessions, subscription.SessionID)
					}
				}
			}

		case message := <-h.Broadcast:
			connections := h.Sessions[message.SessionID]
			for conn := range connections {
				select {
				case conn.Send <- message.Data:
				default:
					close(conn.Send)
					delete(connections, conn)
					if len(connections) == 0 {
						delete(h.Sessions, message.SessionID)
					}
				}
			}
		}
	}
}
