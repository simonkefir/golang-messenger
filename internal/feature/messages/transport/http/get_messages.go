package messages_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// GetUser      godoc
// @Summary     Получить сообщения из чата
// @Description Получить сообщения из чата по его ID
// @Tags        messages
// @Produce     json
// @Param       id    path   int      true                    "ID чата"
// @Security    BearerAuth
// @Success     200 MessageDTOResponse                        "Успешно полученные сообщения"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /chats/{chat_id}/messages                     [get]
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
