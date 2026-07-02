package chats_transport_http

import (
	"net/http"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *ChatsHTTPHandler) GetChats(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	chats, err := h.svc.GetChats(r.Context(), userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get chats")
		return
	}

	response := make([]ChatListItemDTO, 0, len(chats))
	for _, chat := range chats {
		response = append(response, ChatListItemDTO{
			ID:            chat.ID,
			CreatedAt:     chat.CreatedAt.Format("2006-01-02 15:04:05"),
			CompanionID:   chat.CompanionID,
			CompanionName: chat.CompanionName,
		})
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
