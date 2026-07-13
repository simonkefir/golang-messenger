package messages_transport_http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// DeleteMessage godoc
// @Summary      Удаление сообщения в чате
// @Description  Удаление сообщения по ID чата и ID сообщения
// @Tags         messages
// @Security     BearerAuth
// @Success      204                                               "Успешное удаление сообщения в чате"
// @Failure      400 {object} core_http_response.ErrorResponse     "Invalid input"
// @Failure      401 {object} core_http_response.ErrorResponse     "Unauthorized"
// @Failure      403 {object} core_http_response.ErrorResponse     "Forbidden"
// @Failure      404 {object} core_http_response.ErrorResponse     "Not found"
// @Failure      500 {object} core_http_response.ErrorResponse     "Internal server error"
// @Router       /chats/{chat_id}/messages                         [delete]
func (h *MessagesHTTPHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
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
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid chat id")
		return
	}

	var dto DeleteMessageDTO

	if err := core_http_request.DecodeJSON(r, &dto); err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid json")
		return
	}

	if err := h.validate.Struct(dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make(map[string]string)
			for _, fe := range ve {
				errs[fe.Field()] = fe.Tag()
			}
			core_http_response.ValidationError(w, errs)
			return
		}
	}

	if err := h.svc.DeleteMessage(r.Context(), userID, chatID, dto.ID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete message")
		return
	}

	responseHandler.NoContentResponse()
}
