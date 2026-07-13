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

// CreateMessage   godoc
// @Summary        Отправить сообщение
// @Description    Создать новое сообщение в чате по ID чата
// @Tags           messages
// @Accept         json
// @Produce        json
// @Param          request body     CreateMessageDTO     true        "CreateMessage тело запроса"
// @Security       BearerAuth
// @Success        201     {object} MessageDTOResponse               "Успешно созданное сообщение"
// @Failure        400     {object} core_http_response.ErrorResponse "Invalid input"
// @Failure        403     {object} core_http_response.ErrorResponse "Forbidden"
// @Failure        500     {object} core_http_response.ErrorResponse "Internal server error"
// @Router         /chats/{chat_id}/messages                         [post]
func (h *MessagesHTTPHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
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

	var dto CreateMessageDTO

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

	msg, err := h.svc.CreateMessage(r.Context(), userID, chatID, dto.Content)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create message")
		return
	}

	responseHandler.JSONResponse(NewMessageResponseFromDomain(msg), http.StatusCreated)
}
