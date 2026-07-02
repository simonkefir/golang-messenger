package chats_transport_http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *ChatsHTTPHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var dto CreateChatDTO
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

	chat, err := h.svc.CreateChat(r.Context(), userID, dto.ParticipantID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create chat")
		return
	}

	responseHandler.JSONResponse(chat, http.StatusCreated)
}
