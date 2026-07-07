package users_transport_http

import (
	"net/http"
	"strconv"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	idStr := r.PathValue("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidInput, "invalid id")
		return
	}

	user, err := h.svc.GetUser(r.Context(), userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusOK)
}
