package messages_transport_http

import (
	"net/http"

	core_jwt "github.com/simonkefir/golang-messenger/internal/core/jwt"
	core_websocket "github.com/simonkefir/golang-messenger/internal/core/websocket"
)

func (h *MessagesHTTPHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	userID, err := core_jwt.ValidateToken(token)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := core_websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := core_websocket.NewClient(userID, conn)
	h.hub.Register(client)

	go client.WritePump()
	client.ReadPump(func() {
		h.hub.Unregister(client)
	})
}
