package messages_transport_http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/simonkefir/golang-messenger/internal/core/domain"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_request "github.com/simonkefir/golang-messenger/internal/core/transport/http/request"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *MessagesHTTPHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var dto CreateMessageDTO

	if err := core_http_request.DecodeJSON(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "invalid json")
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

	msg, err := h.svc.CreateMessage(r.Context(), domain.Message{
		Content: dto.Content,
	})
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create message")
		return
	}

	responseHandler.JSONResponse(NewMessageResponseFromDomain(&msg), http.StatusCreated)
}
