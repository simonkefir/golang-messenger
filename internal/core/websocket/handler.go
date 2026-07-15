package core_websocket

import (
	"net/http"

	core_jwt "github.com/simonkefir/golang-messenger/internal/core/jwt"
	"go.uber.org/zap"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

func (h *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	userID, err := core_jwt.ValidateToken(token)
	if err != nil {
		h.hub.logger.Warn("unauthorized: %w", zap.Error(err))
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.hub.logger.Error("websocket upgrade failed", zap.Error(err))
		http.Error(w, "websocket upgrade failed", http.StatusInternalServerError)
		return
	}

	client := NewClient(userID, conn)
	h.hub.Register(client)

	go client.WritePump()
	client.ReadPump(func() {
		h.hub.Unregister(client)
	})
}
