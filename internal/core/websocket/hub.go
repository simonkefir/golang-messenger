package core_websocket

import (
	"encoding/json"
	"sync"
)

type Hub struct {
	mu          sync.RWMutex
	connections map[int64][]*Client
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[int64][]*Client),
	}
}

func (h *Hub) Shutdown() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, clients := range h.connections {
		for _, client := range clients {
			close(client.Send)
		}
	}

	h.connections = make(map[int64][]*Client)
}

func (h *Hub) Register(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.connections[client.UserID] = append(h.connections[client.UserID], client)
}

func (h *Hub) Unregister(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	clients := h.connections[client.UserID]
	for i, c := range clients {
		if c == client {
			clients[i] = clients[len(clients)-1]
			h.connections[client.UserID] = clients[:len(clients)-1]
			break
		}
	}

	if len(h.connections[client.UserID]) == 0 {
		delete(h.connections, client.UserID)
	}

	close(client.Send)
}

func (h *Hub) SendEventToUser(userID int64, event Event) {
	payload, err := json.Marshal(event)
	if err != nil {
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.connections[userID] {
		select {
		case client.Send <- payload:
		default:
		}
	}
}
