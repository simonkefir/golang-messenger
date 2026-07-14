package chats_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// GetChat      godoc
// @Summary     Получить чат
// @Description Получить информацию о своём чате из системы чатов
// @Tags        chats
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} ChatListItemDTO                  "Успешно полученный чат"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /chats/{chat_id}                              [get]
func (h *ChatsHTTPHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	idStr := r.PathValue("id")
	chatID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid id")
		return
	}

	chat, err := h.svc.GetChat(r.Context(), userID, chatID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get chat")
		return
	}

	responseHandler.JSONResponse(NewChatResponseFromDomain(chat.Chat, ParticipantDTO(chat.Participant)), http.StatusOK)
}
