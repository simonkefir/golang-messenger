package users_transport_http

import (
	"net/http"

	core_errors "github.com/simonkefir/golang-messenger/internal/core/errors"
	core_logger "github.com/simonkefir/golang-messenger/internal/core/logger"
	core_http_middleware "github.com/simonkefir/golang-messenger/internal/core/transport/http/middleware"
	core_http_response "github.com/simonkefir/golang-messenger/internal/core/transport/http/response"
)

// GetMe        godoc
// @Summary     Получить себя
// @Description Получить информацию о себе из системы пользователей
// @Tags        users
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} UserDTOResponse                  "Успешно полученный пользователь"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users/me                                     [get]
func (h *UsersHTTPHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, ok := core_http_middleware.GetUserID(r.Context())
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	user, err := h.svc.GetMe(r.Context(), userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	responseHandler.JSONResponse(NewUserResponseFromDomain(&user), http.StatusOK)
}
