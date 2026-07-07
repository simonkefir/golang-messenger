package messages_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *MessagesHTTPHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	idStr := r.PathValue("chat_id")
	chatID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid id")
		return
	}

	messages, err := h.svc.GetMessages(r.Context(), userID, chatID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get messages")
		return
	}

	response := make([]*MessageDTOResponse, 0, len(messages))
	for _, msg := range messages {
		response = append(response, NewMessageResponseFromDomain(msg))
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
