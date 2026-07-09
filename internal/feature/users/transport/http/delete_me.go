package users_transport_http

import (
	"net/http"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// DeleteMe     godoc
// @Summary     Удаление своего пользователя
// @Description Удаление себя в системе пользователей
// @Produce     json
// @Tags        users
// @Security    BearerAuth
// @Success     204                                           "Успешное удаление"
// @Failure     400 {object} core_http_response.ErrorResponse "Invalid input"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/me [delete]
func (h *UsersHTTPHandler) DeleteMe(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	if err := h.svc.DeleteMe(r.Context(), userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
